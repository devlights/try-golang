//go:build linux

// Linuxの同じ1 CPUで、通常スレッドとSCHED_RRスレッドの進み方を比較します。
//
// 実験の登場人物は、通常ワーカーと比較対象のprobeです。
// 通常ワーカーは計算を続け、probeは短い計算とsleepを繰り返します。
// NORMAL試験とRT試験でprobeの仕事内容は変えず、OSの優先度だけを変更します。
//
// 読むときは、次の3つを別々の設定として捉えてください。
//   - LockOSThread: goroutineを特定のOSスレッドから移動させません。
//   - CPU affinity: そのOSスレッドが動けるCPUを制限します。
//   - SCHED_RR: そのOSスレッドを、通常スレッドより優先する方式で動かします。
//
// GoのスケジューラがgoroutineをOSスレッド上で実行し、
// LinuxのスケジューラがOSスレッドをCPU上で実行します。
// このデモはLinux側の優先度を比較するもので、処理の完了期限は保証しません。
//
// 公式資料:
// https://pkg.go.dev/runtime#LockOSThread
// https://man7.org/linux/man-pages/man2/sched_setattr.2.html
// https://man7.org/linux/man-pages/man2/sched_setaffinity.2.html
// https://man7.org/linux/man-pages/man7/sched.7.html
package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"runtime"
	"time"

	"golang.org/x/sys/unix"
)

// configは、試験条件をワーカーへ値渡しするための設定です。
// 各ワーカーはコピーを読み取るだけで、実行中に設定を共有更新しません。
type config struct {
	mode     string        // compareは両試験、normal/rrは片方だけ実行します。
	cpu      int           // 競合させる論理CPU番号です。-1はrun内で実番号へ解決します。
	normals  int           // probeとは別に作る通常ワーカーの本数です。
	priority uint          // RT試験のprobeに設定するリアルタイム優先度です。
	nice     int32         // 起動時のnice値を保持し、通常側の条件をそろえます。
	duration time.Duration // 1試験の共通測定窓の長さです。
	busy     time.Duration // probeの1回の計算区間です。CPU時間ではなく壁時計です。
	pause    time.Duration // probeが計算区間の後にsleepする長さです。
}

// validateは、設定値の間違いと、デモとして過大な負荷になる指定を拒否します。
// ここでの上限はLinux自体の制限ではなく、このサンプルの安全寄りの制限です。
func (c config) validate() error {
	if c.mode != "compare" && c.mode != "normal" && c.mode != "rr" {
		return errors.New("modeはcompare、normal、rrのいずれかです")
	}
	if c.cpu < -1 || c.cpu >= 1024 {
		return errors.New("cpuは-1（自動選択）または0～1023です")
	}
	if c.normals < 1 || c.normals > 16 {
		return errors.New("normal-workersは1～16です")
	}
	if c.priority < 1 || c.priority > 20 {
		return errors.New("このデモのpriorityは1～20に制限しています")
	}
	if c.duration < 200*time.Millisecond || c.duration > 10*time.Second {
		return errors.New("durationは200ms～10sです")
	}
	if c.busy < time.Millisecond || c.busy > 20*time.Millisecond {
		return errors.New("busyは1ms～20msです")
	}
	// probeを計算し続けるスレッドにしないため、busy以上のsleepを要求します。
	// 実際のCPU使用率を50%に固定・保証する設定ではありません。
	if c.pause < c.busy || c.pause > time.Second {
		return errors.New("sleepはbusy以上、1s以下にしてください（計算区間は最大50%）")
	}
	// 測定窓が短すぎて、計算とsleepをほとんど試せない指定を避けます。
	if c.duration < 4*(c.busy+c.pause) {
		return errors.New("durationはbusyとsleepの合計の4倍以上にしてください")
	}
	return nil
}

// measurementは、1ワーカーが終了時に返す結果です。
// 実行中はそのワーカーだけが書き、完成後にdoneチャネルで制御側へ渡します。
type measurement struct {
	name     string        // "probe"または"normal-N"です。
	tid      int           // LinuxのOSスレッドIDです。goroutine IDではありません。
	cpu      int           // affinityを検証済みの、固定先CPU番号です。
	policy   uint32        // 要求値ではなく、カーネルから再取得した方式です。
	priority uint32        // 同じく、再取得したリアルタイム優先度です。
	cpuTime  time.Duration // このOSスレッド自身が消費したCPU時間の差分です。
	batches  uint64        // burnを完了した回数です。
	checksum uint64        // 最後の計算結果です。結果を出力まで使うため保持します。
	err      error         // 設定・確認・計測で発生したエラーです。
}

// trialは、NORMALまたはSCHED_RRという1回分の試験結果をまとめます。
type trial struct {
	label    string
	elapsed  time.Duration // 開始合図を送り始めてから、全結果を回収するまでの壁時計時間です。
	duration time.Duration // 指定された測定窓です。CPU使用率の分母に使います。
	rows     []measurement
}

// windowは、全ワーカーへ同じ測定終了時刻を渡すメッセージです。
// time.Now由来のtime.Timeを同じプロセス内で使い、比較には単調時計が利用されます。
type window struct{ end time.Time }

// threadCPUTimeは、呼び出し元OSスレッドがこれまでに消費したCPU時間を取得します。
// 差分を取ると、そのスレッドが実際にCPU上で動いていた時間を測れます。
// sleepやCPUの順番待ちの時間は含みません。goroutine単位の時計ではないので、
// 呼び出す側ではLockOSThreadにより、開始と終了で同じOSスレッドを使います。
func threadCPUTime() (time.Duration, error) {
	var ts unix.Timespec
	if err := unix.ClockGettime(unix.CLOCK_THREAD_CPUTIME_ID, &ts); err != nil {
		return 0, fmt.Errorf("スレッドCPU時計: %w", err)
	}
	return time.Duration(ts.Sec)*time.Second + time.Duration(ts.Nsec), nil
}

// burnは、整数のビット演算を繰り返してCPUを使う、小さな計算単位です。
// 暗号用の乱数生成や有用な業務処理ではなく、比較用の負荷を作るのが目的です。
// 1回を短く区切ることで、呼び出し側は各回の間に測定終了時刻を確認できます。
// 戻り値を次の入力と最後の出力に使い、計算を不要として削除されにくくします。
func burn(x uint64) uint64 {
	for i := 0; i < 4096; i++ {
		x ^= x << 13
		x ^= x >> 7
		x ^= x << 17
	}
	return x
}

// workerは、1本のgoroutineを専用OSスレッドに固定して負荷と計測を行う心臓部です。
//
// index=0がprobeで、それ以外が連続計算する通常ワーカーです。
// rr=trueでもRT化するのはprobeだけです。通常ワーカーの条件は変えません。
//
// 全体の流れ:
//  1. OSスレッドを固定します。
//  2. 優先度とCPU affinityを設定し、カーネルから読み戻して検証します。
//  3. readyで準備結果を通知し、startで共通の終了時刻を受け取ります。
//  4. スレッドCPU時間を測りながら、期限まで計算します。
//  5. doneで結果を返し、OSスレッドをロックしたままgoroutineを終了します。
//
// チャネルの向き:
//   - ready/doneはworkerからrunTrialへの通知専用です。
//   - startはrunTrialからworkerへの開始通知専用です。
//   - 正常なworkerも準備に失敗したworkerも、readyとdoneへ各1回送信します。
func worker(c config, index int, rr bool, start <-chan window, ready chan<- error, done chan<- measurement) {
	// 【1. goroutineとOSスレッドの対応を固定】
	// 通常のgoroutineは、実行再開時に別のOSスレッドへ移ることがあります。
	// Linuxの優先度・affinity・CPU時計はOSスレッド単位なので、
	// 設定するスレッドと、その後に計算・計測するスレッドを同一に保ちます。
	runtime.LockOSThread()
	// 意図的にdefer runtime.UnlockOSThread()は書きません。
	// 設定を戻さずUnlockすると、変更済みのOSスレッドを別のgoroutineへ
	// 再利用させるおそれがあるためです。ロックしたままこのgoroutineが終了すると、
	// 固定していたOSスレッドも終了します。

	// ロック後に取得するため、このTIDは計測中も変わりません。
	r := measurement{name: fmt.Sprintf("normal-%d", index), tid: unix.Gettid()}
	if index == 0 {
		r.name = "probe"
	}
	// 【準備エラーの共通出口】
	// readyだけではなくdoneにも送るのは、runTrialが両方を人数分回収するためです。
	// 準備失敗時はstartを待たず、呼び出し側が直後にreturnします。
	// 両チャネルには人数分のバッファがあるので、まだ回収側が読んでいなくても
	// 各workerの最初の1件を送れます。
	fail := func(err error) {
		r.err = fmt.Errorf("%s: %w", r.name, err)
		ready <- r.err
		done <- r
	}

	// 【2. Linuxのスケジューリング方式を設定】
	// OSスレッドの起動時状態を当てにせず、全workerが自分の設定を明示します。
	// まず全員を通常方式として組み立て、RT試験のprobeだけをSCHED_RRへ変えます。
	// x/sys/unixでは、SCHED_OTHERに相当する定数名はSCHED_NORMALです。
	attr := &unix.SchedAttr{Policy: unix.SCHED_NORMAL, Nice: c.nice}
	if index == 0 && rr {
		attr.Policy = unix.SCHED_RR
		// SCHED_RRでは、大きい値ほど高優先度です。nice値とは別の尺度です。
		attr.Priority = uint32(c.priority)
		// 子タスク作成時に、特権的な優先度をそのまま伝播させないための指定です。
		// 現在のprobe自身のRT設定を解除するフラグではありません。
		attr.Flags = unix.SCHED_FLAG_RESET_ON_FORK
	}
	// 第1引数0は「今これを呼んでいるOSスレッド」です。全スレッドではありません。
	// 第3引数0は、このシステムコールに追加の呼び出しフラグを付けない指定です。
	// attr.Flagsとは別の引数なので、RESET_ON_FORKはattr側へ設定します。
	if err := unix.SchedSetAttr(0, attr, 0); err != nil {
		// 権限不足などを黙って無視すると、「RTのつもりで通常方式を測る」
		// 実験になってしまうため、負荷を開始せずに失敗として返します。
		fail(fmt.Errorf("スケジューリング設定: %w", err))
		return
	}
	// 【3. 実行可能なCPUを1個だけに制限】
	// 全workerが同じCPUを使うことで、本当にそのCPUの取り合いを作ります。
	// 別CPUへ分散すると、優先度の効果ではなく並列実行の差を測りかねません。
	// これは他プロセスをCPUから追い出す「CPUの予約・専有」ではありません。
	var mask unix.CPUSet
	mask.Set(c.cpu)
	// ここでも0は呼び出し元OSスレッドです。ほかのworkerも各自で同じ設定をします。
	if err := unix.SchedSetaffinity(0, &mask); err != nil {
		fail(fmt.Errorf("CPU固定: %w", err))
		return
	}
	// 【4. 実設定を読み戻して検証】
	// 設定要求がエラーにならなかっただけで済ませず、実際の方式・優先度を確認します。
	actual, err := unix.SchedGetAttr(0, 0)
	if err != nil {
		fail(fmt.Errorf("設定確認: %w", err))
		return
	}
	if actual.Policy != attr.Policy || actual.Priority != attr.Priority {
		fail(fmt.Errorf("設定値不一致: policy=%d priority=%d", actual.Policy, actual.Priority))
		return
	}
	var actualMask unix.CPUSet
	if err := unix.SchedGetaffinity(0, &actualMask); err != nil {
		fail(fmt.Errorf("CPU固定確認: %w", err))
		return
	}
	// 目的のCPUを含むだけでは不十分です。「そのCPUだけ」のマスクを要求します。
	if actualMask.Count() != 1 || !actualMask.IsSet(c.cpu) {
		fail(errors.New("指定した1 CPUだけに固定できませんでした"))
		return
	}
	r.policy, r.priority, r.cpu = actual.Policy, actual.Priority, c.cpu
	// 【5. 全員の準備が終わるまで待機】
	// 設定処理の速いworkerだけが先に計算を始めないよう、ここで合図を待ちます。
	// readyへnilを送るのは「結果がない」ではなく「準備成功」という意味です。
	ready <- nil
	w := <-start
	// 全員が同じ終了時刻を受け取りますが、CPUを共有しているので、
	// 全workerがまったく同じ瞬間に実行開始するわけではありません。
	// 準備処理やstart待ちを除外するため、CPU時計は合図を受けてから取得します。
	begin, err := threadCPUTime()
	if err != nil {
		// readyはすでに通知済みなので、この段階のエラーはdoneだけで返します。
		r.err = err
		done <- r
		return
	}
	// 【6. 期限まで負荷を実行】
	// 0以外の初期値を使い、前回の計算結果を次回へ引き継ぎます。
	// 外側のループは、probeでは「計算→sleep」の1サイクルを表します。
	// 通常workerはsleepせず、基本的に内側のループを共通終了時刻まで回します。
	x := uint64(index + 1)
	for time.Now().Before(w.end) {
		// 通常workerの計算期限は、そのまま試験全体の終了時刻です。
		until := w.end
		if index == 0 {
			// probeだけ、今回の計算区間を「今からbusy後」までに短縮します。
			// 試験終了をまたがないよう、共通終了時刻より早い場合だけ採用します。
			//
			// 重要: busyは「CPU時間を5ms使う」ではなく「壁時計で5ms」です。
			// 途中でほかのスレッドにCPUを取られた時間も、この5msに含まれます。
			// そのため、優先される側は同じ壁時計区間でも多く計算できます。
			// NORMAL試験とRT試験で、ここから先の計算方法自体は変えません。
			if t := time.Now().Add(c.busy); t.Before(until) {
				until = t
			}
		}
		// 小さな計算単位ごとに壁時計期限を確認します。
		// burnの途中を強制中断するわけではないため、期限をわずかに超える可能性はあります。
		for time.Now().Before(until) {
			x = burn(x)
			r.batches++
		}
		if index == 0 {
			if remaining := time.Until(w.end); remaining > 0 {
				// probeは明示的に待機し、その間は通常workerがCPUを使えるようにします。
				// LockOSThreadしていてもsleepは可能です。固定が解除されるわけではなく、
				// probeは同じOSスレッド上で処理を再開します。
				//
				// sleepから起床可能になった後、通常workerと競合しているCPUへ
				// どれだけ早く戻れるかが、優先度による差を生むポイントです。
				// ただし、指定時間ちょうどに再開する保証はありません。
				// 最終サイクルでは、要求するsleepを試験の残り時間以下に抑えます。
				time.Sleep(min(c.pause, remaining))
			}
		}
	}
	// 【7. CPU時間の差分と計算結果を返す】
	// 壁時計で2秒の試験でも、CPU時間が0.5秒なら約25%を使ったことになります。
	// 計測区間のCPU時間には、burn以外のループ制御や時計取得なども含まれます。
	finish, err := threadCPUTime()
	r.cpuTime, r.checksum, r.err = finish-begin, x, err
	done <- r
	// ここから関数を抜けるとgoroutineが終了し、ロックしたOSスレッドも終了します。
	// doneの受信は結果が届いた合図であり、OSスレッドの終了完了を待つjoinではありません。
}

// runTrialは、NORMALまたはSCHED_RRの1試験を準備・開始・回収します。
// 負荷はworkerが担当し、ここでは全員を同じ測定窓で動かす段取りを担当します。
func runTrial(c config, rr bool) (trial, error) {
	// 通常workerの本数に、probe 1本を加えます。
	n := c.normals + 1
	// 各workerはready/doneへ1件ずつ送るので、人数分のバッファを持たせます。
	// 準備中に失敗したworkerも、制御側の受信開始を待たずに結果を置けます。
	ready := make(chan error, n)
	done := make(chan measurement, n)
	// startはcloseによる一斉通知ではありません。
	// 同じ終了時刻のメッセージを、workerの人数分送信する方式です。
	start := make(chan window, n)
	t := trial{label: "NORMAL", duration: c.duration, rows: make([]measurement, n)}
	if rr {
		t.label = "SCHED_RR"
	}
	for i := 0; i < n; i++ {
		go worker(c, i, rr, start, ready, done)
	}
	// 全員の準備報告が届くまで測定窓は開始しません。
	// errors.Joinはnilを無視し、複数のエラーがあればまとめて保持します。
	var setupErr error
	for i := 0; i < n; i++ {
		setupErr = errors.Join(setupErr, <-ready)
	}
	now := time.Now()
	end := now.Add(c.duration)
	if setupErr != nil {
		// 準備できたworkerもstartで待っているので、終了済みの窓を渡して解放します。
		// これにより負荷ループに入らず、doneへ結果を返して終了できます。
		end = now // 1本でも準備に失敗したら負荷を開始しません。
	}
	for i := 0; i < n; i++ {
		// 準備失敗で戻ったworkerは受信しませんが、バッファが人数分あるため、
		// その分のメッセージもここへ格納でき、送信待ちで止まりません。
		start <- window{end: end}
	}
	// 成功・失敗にかかわらず全員の結果を回収し、失敗したまま放置しません。
	var workErr error
	for i := 0; i < n; i++ {
		t.rows[i] = <-done
		workErr = errors.Join(workErr, t.rows[i].err)
	}
	// 指定窓durationと、結果回収まで含む実経過elapsedを分けて記録します。
	t.elapsed = time.Since(now)
	if setupErr != nil {
		return t, setupErr
	}
	if workErr != nil {
		return t, workErr
	}
	// ここまでのrowsは結果の到着順です。表示・倍率計算でrows[0]をprobeとして
	// 扱えるよう先頭へ移します。通常worker同士の順番はそろえていません。
	for i := range t.rows {
		if t.rows[i].name == "probe" {
			t.rows[0], t.rows[i] = t.rows[i], t.rows[0]
			break
		}
	}
	return t, nil
}

// printTrialは、回収済みの結果を表示します。workerの計測中には出力しません。
// ログ出力の待ち時間を、負荷ループへ持ち込まないための分離です。
func printTrial(w io.Writer, t trial) {
	fmt.Fprintf(w, "\n[%s] 測定窓=%s 終了回収まで=%s\n", t.label, t.duration, t.elapsed.Round(time.Millisecond))
	fmt.Fprintln(w, "worker       TID    CPU policy       prio CPU-ms  1CPU比%   batches")
	var total time.Duration
	for _, r := range t.rows {
		policy := "SCHED_OTHER"
		if r.policy == unix.SCHED_RR {
			policy = "SCHED_RR"
		}
		fmt.Fprintf(w, "%-10s %7d %4d %-12s %4d %7.1f %8.1f %9d\n",
			r.name, r.tid, r.cpu, policy, r.priority,
			float64(r.cpuTime)/float64(time.Millisecond),
			// 分母は全CPUの時間ではなく、1本のCPUに対する測定窓です。
			// 2秒窓でCPU時間1秒なら、ここでは50%と表示します。
			100*float64(r.cpuTime)/float64(t.duration), r.batches)
		total += r.cpuTime
	}
	if total > 0 {
		// 上の1CPU比%とは分母が異なります。
		// こちらは「測ったworker全体のCPU時間のうちprobeが占めた割合」です。
		// 他プロセスや、別スレッドで動くGoランタイムのCPU時間は含めていません。
		fmt.Fprintf(w, "probeのワーカーCPU時間内シェア: %.1f%%\n", 100*float64(t.rows[0].cpuTime)/float64(total))
	}
	// checksumも参照して計算結果をプログラムの観測可能な出力に含めます。
	fmt.Fprintf(w, "probe checksum: %016x\n", t.rows[0].checksum)
}

// chooseCPUは、競合用CPUと、制御側に許可する残りのCPU集合を決めます。
// 戻り値は「競合CPU番号」「制御側のCPU集合」「エラー」の順です。
// この関数は集合を計算するだけで、OSのaffinity設定自体は変更しません。
func chooseCPU(requested int, allowed unix.CPUSet) (int, unix.CPUSet, error) {
	if allowed.Count() < 2 {
		return 0, allowed, errors.New("制御用CPUを残すため、利用可能なCPUが2個以上必要です")
	}
	target := requested
	if target == -1 {
		// CPU番号は連番で許可されているとは限りません。
		// 0を無条件に使うのではなく、許可マスクから最初のCPUを探します。
		for i := 0; i < 1024; i++ {
			if allowed.IsSet(i) {
				target = i
				break
			}
		}
	}
	if target < 0 || target >= 1024 || !allowed.IsSet(target) {
		return 0, allowed, fmt.Errorf("CPU %dはこのプロセスに許可されていません", target)
	}
	// CPUSetを値コピーし、競合CPUだけを制御側の候補から外します。
	// 元のallowedは変更せず、CPUの専有・予約やシステム設定の変更もしません。
	controller := allowed
	controller.Clear(target)
	return target, controller, nil
}

// runは、引数解析、制御スレッドの配置、試験の実行、結果比較を統括します。
// エラー表示とプロセスの終了コードはmainに任せ、ここではerrorを返します。
func run() error {
	c := config{}
	flag.StringVar(&c.mode, "mode", "compare", "compare / normal / rr")
	flag.IntVar(&c.cpu, "cpu", -1, "競合させるCPU番号。-1は許可CPUから自動選択")
	flag.IntVar(&c.normals, "normal-workers", 3, "通常スレッド数（1～16）")
	flag.UintVar(&c.priority, "priority", 10, "probeのSCHED_RR優先度（1～20）")
	flag.DurationVar(&c.duration, "duration", 2*time.Second, "1試験の測定時間（200ms～10s）")
	flag.DurationVar(&c.busy, "busy", 5*time.Millisecond, "probeが計算する壁時計時間（1ms～20ms）")
	flag.DurationVar(&c.pause, "sleep", 5*time.Millisecond, "probeのsleep時間（busy以上）")
	flag.Parse()
	// --modeなどのフラグ以外に、余分な引数が渡された場合は入力ミスとします。
	if flag.NArg() != 0 {
		return errors.New("位置引数は指定できません")
	}
	if err := c.validate(); err != nil {
		return err
	}
	// PはGoコードを同時に実行するためのランタイム側の枠です。
	// 固定OSスレッドを増やしても、Pが足りなければGo側で順番待ちになります。
	// 今回はLinux側の優先度を見たいので、workerと制御側に余裕のある枠数を設定します。
	// これはCPU数を増やす操作でも、Goの影響を完全に除く操作でもありません。
	runtime.GOMAXPROCS(c.normals + 3)
	runtime.LockOSThread()
	// mainもロックしたままプロセス終了します。affinityを変えたスレッドを返しません。
	// プログラム全体を外からchrtでRT化すると、probeだけを変更する比較が崩れます。
	// 起動時の制御スレッドは通常方式であることを確認します。
	initial, err := unix.SchedGetAttr(0, 0)
	if err != nil {
		return fmt.Errorf("起動時ポリシー取得: %w", err)
	}
	if initial.Policy != unix.SCHED_NORMAL {
		return errors.New("起動時はSCHED_OTHERにしてください。chrtでプログラム全体をRT化しないでください")
	}
	// 通常方式のnice値は起動時の値を全workerへ渡し、同じ条件で比較します。
	c.nice = initial.Nice
	var allowed unix.CPUSet
	// マシン全体のCPU数ではなく、このスレッドに実際に許可された集合を取得します。
	// コンテナやtasksetにより、一部CPUしか使えない場合を考慮しています。
	if err := unix.SchedGetaffinity(0, &allowed); err != nil {
		return fmt.Errorf("許可CPU取得: %w", err)
	}
	target, controller, err := chooseCPU(c.cpu, allowed)
	if err != nil {
		return err
	}
	c.cpu = target
	// 結果回収を行う制御スレッド自身は、競合CPU以外で動けるようにします。
	// ほかの既存ランタイムスレッドや他プロセスまで一括で固定する処理ではありません。
	if err := unix.SchedSetaffinity(0, &controller); err != nil {
		return fmt.Errorf("制御スレッドのCPU分離: %w", err)
	}
	// 実験条件を出力します。RT側のRESET_ON_FORKはworker内で設定しています。
	fmt.Printf("Go=%s 競合CPU=%d 許可CPU数=%d GOMAXPROCS=%d\n",
		runtime.Version(), c.cpu, allowed.Count(), runtime.GOMAXPROCS(0))
	fmt.Printf("normal=%d probe=1 busy=%s sleep=%s（壁時計）通常nice=%d\n",
		c.normals, c.busy, c.pause, c.nice)
	for _, name := range []string{"sched_rt_runtime_us", "sched_rt_period_us"} {
		// RT帯域制限は観測だけです。sysctlへの書き込みや制限の解除は行いません。
		if b, err := os.ReadFile("/proc/sys/kernel/" + name); err == nil {
			fmt.Printf("%s=%s", name, b)
		}
	}
	var base trial
	if c.mode != "rr" {
		// normal/compareでは、まず全員が通常方式の基準試験を行います。
		base, err = runTrial(c, false)
		if err != nil {
			return fmt.Errorf("通常試験: %w", err)
		}
		printTrial(os.Stdout, base)
	}
	if c.mode != "normal" {
		// rr/compareでは、probeだけをRT化した別の試験を行います。
		// runTrialごとに新しいworkerを作るため、基準試験のスレッドを使い回しません。
		rt, err := runTrial(c, true)
		if err != nil {
			return fmt.Errorf("RT試験を実行できません（通常優先度への代替は行いません）: %w\n"+
				"CAP_SYS_NICE、RLIMIT_RTPRIO、コンテナ/cgroupの許可を確認してください", err)
		}
		printTrial(os.Stdout, rt)
		if c.mode == "compare" && base.rows[0].cpuTime > 0 {
			// 測定窓と処理内容は同じです。probeが使えたCPU時間の比を比較します。
			// CPUが高速化した倍率や、必ず得られる性能向上率ではありません。
			// 基準が0のときは、ゼロ除算を避けるため表示しません。
			ratio := float64(rt.rows[0].cpuTime) / float64(base.rows[0].cpuTime)
			fmt.Printf("\nprobe CPU時間比 RT/NORMAL = %.2fx\n", ratio)
			fmt.Println("RTでprobeのCPU時間・シェアが増え、normal側が減るか確認してください。")
			fmt.Println("これは実測値であり、倍率や処理期限の保証ではありません。")
		}
	}
	return nil
}

// mainは、runの結果をプロセスの成否へ変換するだけの入口です。
// エラー時は標準エラーへ説明を出し、呼び出し元シェルへ終了コード1を返します。
func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "エラー:", err)
		os.Exit(1)
	}
}

package signals

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Stop は、signal.Stop()のサンプルです.
//
// signal.Notify()などでハンドル処理を追加した場合
// ファイルの場合と同様に defer で signal.Stop() を呼ぶべき。
//
// > Stop causes package signal to stop relaying incoming signals to c.
// It undoes the effect of all prior calls to Notify using c. When Stop returns, it is guaranteed that c will receive no more signals.
//
// > Stopは、パッケージ・シグナルがcへの受信シグナルのリレーを停止させる。
// ストップが戻れば、cはそれ以上シグナルを受け取らないことが保証される。
//
// # REFERENCES
//   - https://pkg.go.dev/os/signal@go1.24.2#Stop
func Stop() error {
	var (
		l    = log.New(os.Stderr, "", log.Lmicroseconds)
		pid  = os.Getpid()
		sigs = make(chan os.Signal, 1)

		sendSig = func() {
			l.Println(">> SIGUSR1")
			_ = syscall.Kill(pid, syscall.SIGUSR1)
		}
		stopSig = func() {
			l.Println(">> signal.Stop")
			signal.Stop(sigs)
		}
	)
	defer close(sigs)

	signal.Notify(sigs, syscall.SIGUSR1, syscall.SIGUSR2)
	go func() {
		l.Printf("Signal recv: %v\n", <-sigs)
	}()

	l.Println("<START>")
	{
		sendSig()
		stopSig()
		sendSig()
	}
	time.Sleep(1 * time.Second)
	l.Println("< END >")

	/*
		$ task
		task: [build] go build -o "/workspaces/try-golang/try-golang" .
		task: [run] ./try-golang -onetime

		ENTER EXAMPLE NAME: signal_stop

		[Name] "signal_stop"
		06:03:32.489167 <START>
		06:03:32.489314 >> SIGUSR1
		06:03:32.489326 >> signal.Stop
		06:03:32.489444 >> SIGUSR1
		06:03:32.489442 Signal recv: user defined signal 1
		06:03:33.489539 < END >


		[Elapsed] 1.000858245s
	*/

	return nil
}

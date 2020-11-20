package goroutines

import (
	"context"
	"log"
	"os"
	"time"
)

// SelectNilChan1 -- select ステートメントで nil チャネル を使って選択されるチャネルの有効・無効を切り替えるサンプルです (1).
func SelectNilChan1() error {
	// ２つのゴルーチンを起動
	//   - 一つの生存期間は　　 2 秒
	//   - もう一つの生存期間は 5 秒
	// メインゴルーチンで待ち合わせをして、終わり次第結果を報告する select を一つ用意して待つ
	// メイン処理の生存期間は 7 秒とする

	// select ステートメントで case に指定するチャネルは
	// その値が nil の場合は決して選択されない。（ nil チャネルは送信も受信も不可のため）
	// これを利用すると、case に指定しているチャネルの有効・無効を切り替えることが出来る
	// 無効にしたい場合は nil にして、有効にしたい場合は nil 以外の値を入れる

	// 出力用ロガー
	var (
		mainLog = log.New(os.Stdout, "[main] ", 0)
		g1Log   = log.New(os.Stderr, ">>> G1 ", 0)
		g2Log   = log.New(os.Stderr, ">>> G2 ", 0)
	)

	// 生存期間
	var (
		g1Timeout   = 2 * time.Second
		g2Timeout   = 5 * time.Second
		procTimeout = 7 * time.Second
	)

	// コンテキスト定義
	var (
		rootCtx             = context.Background()
		mainCtx, mainCancel = context.WithCancel(rootCtx)
		procCtx, procCancel = context.WithTimeout(mainCtx, procTimeout)
	)

	defer mainCancel()
	defer procCancel()

	mainLog.Println("start", time.Now().UTC().Unix())

	// ゴルーチン起動
	g1Ctx := g(procCtx, g1Timeout, g1Log)
	g2Ctx := g(procCtx, g2Timeout, g2Log)

	// 待ち合わせ
	var (
		doneProc, doneG1, doneG2 = procCtx.Done(), g1Ctx.Done(), g2Ctx.Done()
	)

LOOP:
	for {
		select {
		case <-doneG1:
			mainLog.Println("g1 done", time.Now().UTC().Unix())

			// このチャネルはクローズされているので、そのままにしておくと
			// 永遠と選択候補として残ってしまう。次から無効とするために、nil にする.
			// これにより、次から選択されなくなる.(selectの選択対象とならない)
			doneG1 = nil
		case <-doneG2:
			mainLog.Println("g2 done", time.Now().UTC().Unix())
			doneG2 = nil
		case <-doneProc:
			mainLog.Println("proc done", time.Now().UTC().Unix())
			break LOOP
		}
	}

	return nil
}

func g(pCtx context.Context, timeout time.Duration, logger *log.Logger) context.Context {
	ctx, cancel := context.WithTimeout(pCtx, timeout)

	go func() {
		defer cancel()

	LOOP:
		for {
			select {
			case <-ctx.Done():
				logger.Println("done", time.Now().UTC().Unix())
				break LOOP
			default:
			}

			logger.Println("processing...", time.Now().UTC().Unix())

			select {
			case <-ctx.Done():
			case <-time.After(1 * time.Second):
			}
		}
	}()

	return ctx
}

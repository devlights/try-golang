package goroutines

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/devlights/gomy/chans"
)

// SelectNilChan2 -- select ステートメントで nil チャネル を使って選択されるチャネルの有効・無効を切り替えるサンプルです (2).
func SelectNilChan2() error {
	// ２つのゴルーチンを起動し、５秒毎に出力を切り替え
	// メインゴルーチンで待ち合わせをして、進捗を報告する select を一つ用意して待つ
	// メイン処理の生存期間は 12 秒、起動するゴルーチンの生存期間は 10 秒とする

	// select ステートメントで case に指定するチャネルは
	// その値が nil の場合は決して選択されない。（ nil チャネルは送信も受信も不可のため）
	// これを利用すると、case に指定しているチャネルの有効・無効を切り替えることが出来る
	// 無効にしたい場合は nil にして、有効にしたい場合は nil 以外の値を入れる
	//
	// また、チャネル自体の値を入れ替えることで別のチャネルに切り替えることも可能

	// ロガー
	//
	// 有効にしたい場合は、ioutil.Discard を 望みの io.Writer に変更
	var (
		g1Log      = log.New(ioutil.Discard, ">>> [G1] ", 0)
		g2Log      = log.New(ioutil.Discard, ">>> [G2] ", 0)
		monitorLog = log.New(ioutil.Discard, ">>>>> [monitor] ", 0)
		mainLog    = log.New(os.Stdout, "[main] ", 0)
	)

	// 生存期間
	var (
		g1Timeout   = 10 * time.Second
		g2Timeout   = 10 * time.Second
		procTimeout = 12 * time.Second
	)

	// チャネル
	var (
		monitorCh  = make(chan chan string)
		g1StatusCh = make(chan string)
		g2StatusCh = make(chan string)
	)

	// コンテキスト
	var (
		rootCtx             = context.Background()
		mainCtx, mainCancel = context.WithCancel(rootCtx)
		procCtx, procCancel = context.WithTimeout(mainCtx, procTimeout)
	)

	defer mainCancel()
	defer procCancel()

	m := startMonitor(procCtx, g1StatusCh, g2StatusCh, monitorCh, monitorLog)
	g1 := startG(procCtx, "G1", g1Timeout, g1StatusCh, g1Log)
	g2 := startG(procCtx, "G2", g2Timeout, g2StatusCh, g2Log)

	// 待ち合わせながら状況を出力
	var (
		statusCh <-chan string
	)

LOOP:
	for {
		select {
		case <-procCtx.Done():
			mainLog.Println("proc done", time.Now().UTC().Unix())
			break LOOP
		case s, ok := <-monitorCh:
			if !ok {
				break LOOP
			}

			// 状況出力するチャネルを切り替え
			statusCh = s
		case v, ok := <-statusCh:
			if !ok {
				break LOOP
			}

			mainLog.Println(v)
		}
	}

	<-chans.WhenAll(m.Done(), g1.Done(), g2.Done())

	return nil
}

func startG(pCtx context.Context, name string, timeout time.Duration, statusCh chan string, l *log.Logger) context.Context {
	ctx, cancel := context.WithTimeout(pCtx, timeout)

	go func() {
		defer cancel()

		for {
			select {
			case <-ctx.Done():
				l.Println("done", time.Now().UTC().Unix())
				return
			case statusCh <- fmt.Sprintf("[%s] running... \t%d", name, time.Now().UTC().Unix()):
			}

			l.Println(time.Now().UTC().Unix())

			select {
			case <-ctx.Done():
			case <-time.After(1 * time.Second):
			}
		}
	}()

	return ctx
}

func startMonitor(pCtx context.Context, g1, g2 chan string, m chan chan string, l *log.Logger) context.Context {
	ctx, cancel := context.WithCancel(pCtx)

	go func() {
		defer cancel()

		// チャネル定義
		var (
			current  chan string // 現在処理中
			draining chan string // 吸い出し中
		)

		// コンテキスト
		var (
			drainCtx    context.Context
			drainCancel context.CancelFunc
		)

		for {
			select {
			case <-ctx.Done():
				if drainCancel != nil {
					drainCancel()
				}

				l.Println("done", time.Now().UTC().Unix())
				return
			default:
			}

			var name string
			switch current {
			case nil:
				// 本来はここでも ctx.Done() を確認するべきだが 割愛
				m <- g1
				current = g1
				draining = g2
				name = "g2"

				drainCtx, drainCancel = context.WithCancel(pCtx)

				l.Println("prev: none\tcurrent: g1\tdraining: g2")
			case g1:
				m <- g2
				current = g2
				draining = g1
				name = "g1"

				drainCancel()
				drainCtx, drainCancel = context.WithCancel(pCtx)

				l.Println("prev: g1\tcurrent: g2\tdraining: g1")
			case g2:
				m <- g1
				current = g1
				draining = g2
				name = "g2"

				drainCancel()
				drainCtx, drainCancel = context.WithCancel(pCtx)

				l.Println("prev: g2\tcurrent: g1\tdraining: g2")
			}

			// 逆側のゴルーチンの出力を吸い出し
			drain(drainCtx, name, draining, l)

			select {
			case <-ctx.Done():
			case <-time.After(5 * time.Second):
			}
		}
	}()

	return ctx
}

func drain(pCtx context.Context, name string, ch <-chan string, l *log.Logger) {
	ctx, cancel := context.WithCancel(pCtx)
	go func() {
		defer cancel()
		for {
			select {
			case <-ctx.Done():
				l.Printf("drain %s stop\t%d", name, time.Now().UTC().Unix())
				return
			case <-ch:
				l.Printf("<<< draining %s", name)
			}
		}
	}()
}

// OVERVIEW
//
// JAVA VM は SIGQUIT を受け取るとスレッドダンプを出力します。
// Goで同じような動きをするサンプルです。
//
// REFERENCE
//
//   - https://stackoverflow.com/a/27398062

package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

const (
	Interval  = 1 * time.Second
	TimeLimit = 10 * time.Second
)

var (
	appLog  = log.New(os.Stdout, "", 0)
	tickLog = log.New(os.Stdout, "", log.Ltime)
	dumpLog = log.New(os.Stderr, "", 0)
)

func main() {
	// ------------------------------------------------
	// Overview
	// ------------------------------------------------
	// [x] 10秒たったらプログラム終了
	// [x] SIGINTを受けたらプログラム終了
	// [x] SIGQUITを受けたらスレッドダンプを出力
	// [x] 1秒毎の経過ログを出力
	// ------------------------------------------------

	var (
		rootCtx          = context.Background()
		mainCtx, mainCxl = context.WithTimeout(rootCtx, TimeLimit)
	)
	defer mainCxl()

	var (
		sigIntCh  = make(chan os.Signal, 1)
		sigQuitCh = make(chan os.Signal, 1)
	)
	defer close(sigIntCh)
	defer close(sigQuitCh)

	signal.Notify(sigIntCh, syscall.SIGINT)
	signal.Notify(sigQuitCh, syscall.SIGQUIT)

	var (
		ticker = time.NewTicker(Interval)
		count  = 0
	)
	defer ticker.Stop()

LOOP:
	for {
		select {
		case <-mainCtx.Done():
			appLog.Println("Timed out")
			break LOOP
		case <-sigIntCh:
			appLog.Println(" >>> Recv SIGINT")
			break LOOP
		case <-sigQuitCh:
			buf := make([]byte, 1<<25)
			dumpLog.Printf(" >>> *** DUMP ***\n\n%s\n", buf[:runtime.Stack(buf, true)])
		case <-ticker.C:
			count++
			tickLog.Println(count)
		}
	}

	appLog.Println("DONE")
}

//go:build !windows

// シグナルを受信する側です
//
// # 処理手順
//
//   - signal.NotifyContext を利用して SIGTERM をフック
//   - シグナルが送信されるのを待機
//
// REFERENCES:
//   - https://stackoverflow.com/questions/9030680/list-of-currently-running-process-in-go
package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	appLog = log.New(os.Stderr, "[receiver] >>> ", 0)
)

func main() {
	var (
		mainCtx              = context.Background()
		procCtx, procCxl     = context.WithTimeout(mainCtx, 10*time.Second)
		signalCtx, signalCxl = signal.NotifyContext(procCtx, syscall.SIGTERM)
	)
	defer procCxl()
	defer signalCxl()

	appLog.Println("wait for SIGTERM")
	select {
	case <-procCtx.Done():
		appLog.Println("timeout")
	case <-signalCtx.Done():
		appLog.Println("receive SIGTERM from sender")
	}
}

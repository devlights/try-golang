// signal.Reset() のサンプルです。
//
// # REFERENCES
//   - https://pkg.go.dev/os/signal@go1.19.3#Reset
package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	ctx, cxl := context.WithCancel(context.Background())
	defer cxl()

	// 0. 待機用
	waitCtx, waitCxl := context.WithTimeout(ctx, 5*time.Second)
	defer waitCxl()

	// 1. SIGINT を Notify しておく
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)

	// 2. SIGINT のハンドリングと signal.Reset() の呼び出し
	fmt.Println("5秒間 SIGINT をハンドリング")
	go func() {
	LOOP:
		for {
			select {
			case <-waitCtx.Done():
				fmt.Fprintln(os.Stderr, "signal.Reset() CALLED")
				signal.Reset(os.Interrupt)
				break LOOP
			case <-sigCh:
				fmt.Fprintln(os.Stderr, "CTRL-C PRESSED")
			}
		}
	}()

	// 3. signal.Reset() 後の挙動確認 (CTRL-C 押下で元の動き(アプリ終了)に戻っていることの確認)
	<-waitCtx.Done()
	<-time.After(5 * time.Second)
}

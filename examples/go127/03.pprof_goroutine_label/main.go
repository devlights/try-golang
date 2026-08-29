//go:build unix

// runtime/pprof のラベルが Go 1.27 では goroutine のトレースバックにも
// 表示されることを確認するサンプルです。
//
// pprof.Do の動的スコープ内で実行される goroutine のスタック表示へ、pprof ラベルが反映されるようになりました。
//
// pprof.Do で ID ラベルを付与した goroutine 内から runtime.Stack を実行し、
// スタックトレースの先頭に {id: "..."} が出力されることを確認します。
// これにより、障害調査時に goroutine をリクエスト ID やジョブ ID などの
// 処理コンテキストと対応付けやすくなります。
//
// # REFERENCES
//   - https://victoriametrics.com/blog/go-1-27/#goroutine-labels-in-tracebacks
package main

import (
	"context"
	"fmt"
	"runtime"
	"runtime/pprof"
	"sync"
	"uuid"
)

func main() {
	var (
		ctx = context.Background()
		id  = uuid.NewV7()
		lbl = pprof.Labels("id", id.String())
		wg  sync.WaitGroup
	)
	for range 2 {
		wg.Go(func() {
			pprof.Do(ctx, lbl, func(ctx context.Context) {
				var (
					buf = make([]byte, 1<<12)
					n   = runtime.Stack(buf, false)
				)
				fmt.Printf("%s\n", buf[:n])
			})
		})
	}

	wg.Wait()
}

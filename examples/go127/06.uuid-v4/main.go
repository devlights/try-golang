package main

import (
	"context"
	"fmt"
	"os/signal"
	"sync"
	"syscall"

	// Go 1.27 で標準ライブラリに追加された uuid パッケージ。
	// RFC 9562 に準拠し、乱数部分には暗号論的に安全な乱数生成器が使われる。
	// import パスは "crypto/uuid" ではなく "uuid" である点に注意。
	"uuid"
)

func main() {
	var (
		rootCtx       = context.Background()
		mainCtx, stop = signal.NotifyContext(rootCtx, syscall.SIGINT, syscall.SIGTERM)
	)
	defer stop()

	if err := run(mainCtx); err != nil {
		panic(err)
	}
}

func run(_ context.Context) error {
	//
	// Go 1.27 から、uuid パッケージが追加された。
	//
	// v4とv7がサポートされている。
	//
	// - UUID v4: ランダム
	// - UUID v7: 先頭48ビットにタイムスタンプが設定される
	//
	// 補足:
	//   - uuid.New() は「用途を問わない標準の生成関数」で、現時点では NewV4 と同等。
	//     ただし将来アルゴリズムが変わる可能性がドキュメントに明記されている。
	//     バージョンを固定したい場合は NewV4 / NewV7 を明示的に使うこと。
	//   - v7 は DB の主キーに向く。時刻順に並ぶため B-tree への挿入が局所化する。
	//     一方で生成時刻が ID から読み取れてしまうため、外部公開 ID には v4 が無難。
	//
	const (
		Count = 3
	)
	var (
		ch = make(chan uuid.UUID, Count)
		wg sync.WaitGroup
	)
	for range Count {
		wg.Go(func() {
			ch <- uuid.NewV4()
		})
	}
	go func() {
		defer close(ch)
		wg.Wait()
	}()

	for v := range ch {
		fmt.Printf("uuid(v4)=%v\n", v)
	}

	return nil
}

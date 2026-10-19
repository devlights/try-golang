package main

import (
	"context"
	"encoding/binary"
	"fmt"
	"os/signal"
	"sync"
	"syscall"
	"time"

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
		Layout = "15:04:05.000"
	)
	fmt.Printf("now=%s\n", time.Now().Format(Layout))

	const (
		Count = 3
	)
	var (
		ch = make(chan uuid.UUID, Count)
		wg sync.WaitGroup
	)
	for range Count {
		wg.Go(func() {
			ch <- uuid.NewV7()
		})
	}
	go func() {
		defer close(ch)
		wg.Wait()
	}()

	for v := range ch {
		//
		// 実行すると以下のように出力される
		//
		// 01a15335-12f3-739c-8b56-6ab8ed992f55
		//
		// 01a15335-12f3 の部分までが先頭 48ビット（6バイト）となる。(48ビットの符号なし整数)
		// Unix epoch (1970-01-01 00:00:00 UTC からの経過ミリ秒数) がビッグエンディアンで格納されており
		// 後ろ 16ビット にver(4bit)とrand_a(12bit)が付いている。
		//
		// [ Unix ミリ秒: 48ビット ][ version: 4ビット ][ rand_a: 12ビット ]
		//
		// uuid.UUID は [16]byte となっているため、タイムスタンプ部分を抜き出すには以下のように出来る。
		//
		var (
			ts = v[:8]                             // 先頭8バイトをスライス
			ms = binary.BigEndian.Uint64(ts) >> 16 // version と rand_a の合計16ビットを取り除く
			t  = time.UnixMilli(int64(ms))         // unixミリ秒として復元
		)
		fmt.Printf("uuid(v7)=%v\ttimestamp=%s\n", v, t.Format(Layout))
	}

	return nil
}

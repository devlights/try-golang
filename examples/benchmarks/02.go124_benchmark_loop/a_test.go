package main

import (
	"testing"
	"time"
)

func setup() {
	time.Sleep(1 * time.Second)
}

func slowFn() {
	time.Sleep(100 * time.Millisecond)
}

func Benchmark_OldStyleLoop_NoResetTimer(b *testing.B) {
	setup()

	for range b.N {
		slowFn()
	}

	b.Logf("N: %d", b.N)
}

func Benchmark_OldStyleLoop_WithResetTimer(b *testing.B) {
	setup()
	b.ResetTimer() // セットアップの時間を含めないようここでタイマーをリセット

	for range b.N {
		slowFn()
	}

	b.Logf("N: %d", b.N)
}

func Benchmark_NewStyleLoop(b *testing.B) {
	//
	// Go 1.24 にて、testing.B.Loop が追加された。
	// シグネチャは以下の様になっている。
	//
	// 	func (b *B) Loop() bool
	//
	// 以前までは、testing.B.N を使用してループすることで
	// ベンチマークを計測していたが、今後は testing.B.Loop を
	// 使用してベンチマークすることが推奨されるようになる。
	//
	// testing.B.Loop を使用することで以下の恩恵がある。
	//
	// 	- ベンチマーク関数が１回のみ実行されるようになる
	// 		- セットアップやクリーンアップの回数が減少
	// 	- b.Loop を利用しているループはコンパイラの最適化がかからないようになる
	// 	- b.Loop の開始と終了時にタイマーが自動管理されるようになる
	// 		- b.ResetTimer() の呼び出しが不要となる
	// 		- セットアップコードがベンチマーク時間に含まれなくなる
	//
	// # REFERENCES
	// 	- https://pkg.go.dev/testing@go1.24.0#hdr-b_N_style_benchmarks
	// 	- https://pkg.go.dev/testing@go1.24.0#B.Loop
	// 	- https://antonz.org/go-1-24/#benchmark-loop
	// 	- https://www.bytesizego.com/blog/go-124-new-benchmark-function
	//

	// セットアップコードは１度だけ呼び出される
	setup()

	// b.Loop を利用する場合、内部でタイマーの自動管理が行われるので
	// b.ResetTimer() の呼び出しが不要となる
	//
	// - b.Loop の開始時に b.ResetTimer() される
	// - b.Loop が false を返したときに b.StopTimer() される
	for b.Loop() {
		//
		// for b.Loop() { ... } の中は決して最適化が行われない
		//
		slowFn()
	}

	b.Logf("N: %d", b.N)
}

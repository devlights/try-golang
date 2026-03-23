package main

import (
	"context"
	"flag"
	"os"
	"runtime/pprof"
	"time"
)

type (
	_Args struct {
		withprof bool
	}
)

var (
	args _Args
)

func init() {
	flag.BoolVar(&args.withprof, "prof", false, "with pprof")
}

func main() {
	flag.Parse()

	var (
		ctx = context.Background()
		err error
	)
	if args.withprof {
		//
		// Go 1.26 時点では goroutineleak プロファイル利用は実験的要素。
		// GOEXPERIMENT=goroutineleakprofileの指定が必要となる。
		//
		//	https://go.dev/doc/go1.26#goroutineleak-profiles
		//	https://antonz.org/go-1-26/#pprof-goroutineleak
		//

		var (
			prof = pprof.Lookup("goroutineleak")
		)
		defer func() {
			// WriteToメソッドの第二引数
			//
			// 	The debug parameter enables additional output.
			// 	Passing debug=0 writes the gzip-compressed protocol buffer described in https://github.com/google/pprof/tree/main/proto#overview.
			// 	Passing debug=1 writes the legacy text format with comments translating addresses to function names and line numbers, so that a programmer can read the profile without tools.
			//	The predefined profiles may assign meaning to other debug values; for example,
			// 		when printing the "goroutine" profile, debug=2 means to print the goroutine stacks in the same form that a Go program uses when dying due to an unrecovered panic.
			//
			// 本サンプルでは goroutine のプロファイルが見たいので 2 を渡している
			//
			// 事前定義されているプロファイルは以下の名前で登録されている。(go/src/runtime/pprof/pprof.go)
			//
			//	goroutine      - stack traces of all current goroutines
			//	goroutineleak  - stack traces of all leaked goroutines
			//	allocs         - a sampling of all past memory allocations
			//	heap           - a sampling of memory allocations of live objects
			//	threadcreate   - stack traces that led to the creation of new OS threads
			//	block          - stack traces that led to blocking on synchronization primitives
			//	mutex          - stack traces of holders of contended mutexes
			//
			prof.WriteTo(os.Stdout, 2)
		}()

		err = run(ctx)
	} else {
		err = run(ctx)
	}

	if err != nil {
		panic(err)
	}
}

func run(pCtx context.Context) error {
	var (
		ctx, cxl = context.WithTimeout(pCtx, 1*time.Second)
	)
	defer cxl()

	leak() // 戻り値のチャネルを利用していないためゴルーチンがリークする
	<-ctx.Done()

	return nil
}

func leak() <-chan int {
	var (
		out = make(chan int)
	)
	go func() {
		out <- 0xBEEF
	}()

	return out
}

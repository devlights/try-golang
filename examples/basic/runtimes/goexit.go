package runtimes

import (
	"runtime"

	"github.com/devlights/gomy/output"
)

// Goexit -- runtime.Goexit() のサンプルです。
//
// # REFERENCES:
//   - https://dev.to/freakynit/the-very-useful-runtime-package-in-golang-5b16
func Goexit() error {
	//
	// runtime.Goexit() は、呼び出すとカレントコンテキストのgoroutineにて
	// 遅延実行、つまり、deferされているものを実行してから即座にgoroutineを終了させる。
	//
	// 注意点として、この関数をメインゴルーチン上で実行してはいけない。
	//

	var (
		fn1 = func() {
			output.Stdoutl("[fn1]", "hello")
		}
		fn2 = func() {
			output.Stdoutl("[fn2]", "world")
		}
	)

	done := make(chan struct{})
	go func() {
		defer close(done)
		defer func() { fn2() }()
		defer func() { fn1() }()

		output.Stdoutl("[before]", "runtime.Goexit() 呼び出し前")
		runtime.Goexit()
		output.Stdoutl("[after ]", "runtime.Goexit() 呼び出し後 ここには到達しない")
	}()
	<-done

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: runtime_goexit

	   [Name] "runtime_goexit"
	   [before]             runtime.Goexit() 呼び出し前
	   [fn1]                hello
	   [fn2]                world


	   [Elapsed] 94.7µs
	*/

}

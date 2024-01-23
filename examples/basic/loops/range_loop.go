package loops

import "github.com/devlights/gomy/output"

// RangeLoop は、単純に指定回数ループするためのサンプルです.
func RangeLoop() error {
	// 単純に指定回数だけループしたい場合、[]struct{} を作って
	// ループさせるのが効率が良い. struct{} はメモリを消費しない.
	for range make([]struct{}, 3) {
		output.Stdoutl("", "hello")
	}

	// インデックスが欲しい場合
	for i := range make([]struct{}, 3) {
		output.Stdoutf("", "[%d] hello\n", i)
	}

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: loops_range_loop

	   [Name] "loops_range_loop"
	   hello
	   hello
	   hello
	   [0] hello
	   [1] hello
	   [2] hello


	   [Elapsed] 28.85µs
	*/

}

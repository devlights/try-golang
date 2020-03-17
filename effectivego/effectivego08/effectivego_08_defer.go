package effectivego08

import "fmt"

// Defer -- Effective Go - Defer の 内容についてのサンプルです。
func Defer() error {
	/*
		https://golang.org/doc/effective_go.html#defer

		- defer ステートメントは、関数呼び出しを遅延評価する仕組み
		- リソースの後処理などによく利用される (I/O Close, Channel Close, etc,,,)
		- 他言語にある try-finally の仕組みと同じイメージだが、呼び出しがすぐ近くにあるので分かりやすい
		- defer される関数に指定する引数は 実際に遅延評価されて呼ばれる タイミングではなく、deferが評価されたタイミングで評価される
	*/
	for i := 0; i < 5; i++ {
		//noinspection GoDeferInLoop
		defer fmt.Println(i)
	}

	enter := func(message string) string {
		fmt.Println("enter", message)
		return message
	}

	leave := func(message string) {
		fmt.Println("leave", message)
	}

	f1 := func() {
		defer leave(enter("f1"))
		fmt.Println("in f1")
	}

	defer leave(enter("main"))
	f1()

	return nil
}

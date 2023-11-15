package builtins

import "fmt"

// PrintFunc は、ビルドイン関数のprintとfmt.Printの違いについてのサンプルです.
func PrintFunc() error {
	// ------------------------------------------------------------
	// ビルドイン関数の print() と fmt.Print() の違い
	//
	// ビルドイン関数の print(), println() は標準エラーに出力する
	// fmt.PrintXX は、標準出力に出力する
	//
	// ビルトイン関数の方は、fmtパッケージをimportする必要がないため
	// アプリのブート時やデバッグ目的に利用すると便利。
	//
	// ただし、ビルドイン関数である print(), println() のコメントには
	// 以下の一文が記載されている。(builtin/builtin.goより)
	//   "it is not guaranteed to stay in the language."
	// ------------------------------------------------------------
	var (
		message = "helloworld"
	)

	// fmtパッケージの方は標準出力に出力する
	fmt.Println("[fmt    ]", message)

	// ビルトインの方は標準エラーに出力する
	println("[builtin]", message)

	return nil

	/*
	   $ task
	   task: Task "build" is up to date
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: builtin_print

	   [Name] "builtin_print"
	   [fmt    ] helloworld
	   [builtin] helloworld


	   [Elapsed] 7.68µs
	*/

}

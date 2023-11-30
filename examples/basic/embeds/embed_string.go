package embeds

import (
	_ "embed" // Go1.16 の embed パッケージを利用するために必要な import

	"github.com/devlights/gomy/output"
)

//go:embed helloworld.txt
var s string // helloworld.txt の中身が string として設定される

// EmbedString は、embed パッケージの機能を確認するサンプルです (文字列として値を取得)
func EmbedString() error {
	//
	// go:embed で指定されている部分は、コンパイル時に埋め込みが施される
	// なので、プログラム実行時には既に値は変数に設定済みとなっている.
	//
	output.Stdoutl("[helloworld.txt]", s)
	return nil

	/*
	   $ task
	   task: Task "build" is up to date
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: embed_string

	   [Name] "embed_string"
	   [helloworld.txt]     hello
	   world


	   [Elapsed] 6.04µs
	*/

}

package effectivego

import "github.com/devlights/try-golang/lib/output"

// Effective Go - Printing の 内容についてのサンプルです。
func Printing() error {
	/*
			https://golang.org/doc/effective_go.html#printing

			- fmt.Stringer インターフェースを実装していると String() が呼ばれるようになる
			- 書式部分に利用する verbs は以下を参照
		      - https://golang.org/pkg/fmt/#hdr-Printing
	*/
	// %d は10進数で出力
	output.Stdoutf("(1)", "%d\n", 255)
	// %x は16進数で出力
	output.Stdoutf("(2)", "%x\n", 255)
	// %s は文字列で出力
	output.Stdoutf("(3)", "%s\n", "helloworld")
	// %q はダブルクォート付きで文字列出力
	output.Stdoutf("(4)", "%q\n", "helloworld")
	// %v は汎用的に利用できる書式 (value の意味)
	output.Stdoutf("(5)", "%v\n", map[string]string{"hello": "world"})

	type S struct {
		V1 string
		V2 string
	}

	s := S{
		V1: "hello",
		V2: "world",
	}

	// %+v は %v に加えて型の構造も出力してくれる
	output.Stdoutf("(6)", "%+v\n", s)
	// %#v は %+v の出力を Go のシンタックスで出力してくれる
	// v系の中で、これが一番詳細に出力する.
	output.Stdoutf("(7)", "%#v\n", s)
	// %T は 型を出力する
	output.Stdoutf("(8)", "%T\n", s)

	return nil
}

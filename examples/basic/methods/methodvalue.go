package methods

import (
	"bytes"
	"io"
	"os"
)

// MethodValue は、Goのメソッド値のサンプルです.
//
// メソッド値と対をなすものがメソッド式となる。
// メソッド式のサンプルは methodexpression.go を参照。
//
// # REFERENCES
//   - https://go.dev/ref/spec#Method_expressions
//   - https://zenn.dev/spiegel/articles/20201212-method-value-and-expression
//   - https://stackoverflow.com/questions/48883754/can-any-one-explain-this-go-program-which-uses-a-method-expression
func MethodValue() error {
	//
	// メソッド値は、インスタンスが紐づいた状態の関数オブジェクトのこと
	// 既にレシーバーが紐づいているので、呼び出す際にレシーバーを渡す必要はない
	//

	var (
		buf  = &bytes.Buffer{}
		fn   = buf.Write // メソッド値
		data = []byte("hello world")
	)

	if _, err := fn(data); err != nil {
		return err
	}

	io.Copy(os.Stdout, buf)

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: methods_method_value

	   [Name] "methods_method_value"
	   hello world

	   [Elapsed] 7.22µs
	*/
}

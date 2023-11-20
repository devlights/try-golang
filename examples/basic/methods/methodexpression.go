package methods

import (
	"bytes"
	"io"
	"os"
)

// MethodExpression は、Goのメソッド式のサンプルです.
//
// メソッド式と対をなすものがメソッド値となる。
// メソッド値のサンプルは methodvalue.go を参照。
//
// # REFERENCES
//   - https://go.dev/ref/spec#Method_expressions
//   - https://zenn.dev/spiegel/articles/20201212-method-value-and-expression
//   - https://stackoverflow.com/questions/48883754/can-any-one-explain-this-go-program-which-uses-a-method-expression
func MethodExpression() error {
	//
	// メソッド式は、インスタンスが紐づいていない状態の関数オブジェクトのこと
	// レシーバーが紐づいていないので、呼び出す際にレシーバーを渡す必要がある
	// (Python の self と思うと分かりやすい)
	//

	var (
		buf  = &bytes.Buffer{}
		fn   = (*bytes.Buffer).Write // メソッド式
		data = []byte("hello world")
	)

	// メソッド式は、レシーバーが紐づいていないので、第一引数にレシーバーを渡して呼び出す
	if _, err := fn(buf, data); err != nil {
		return err
	}

	io.Copy(os.Stdout, buf)

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: methods_method_expression

	   [Name] "methods_method_expression"
	   hello world

	   [Elapsed] 6.16µs
	*/

}

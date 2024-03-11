package strs

import (
	"strings"

	"github.com/devlights/gomy/output"
)

// TrimSpace は、strings.TrimSpace() のサンプルです.
//
// 両端のスペースをトリミングしてくれる。
//
// > TrimSpace returns a slice of the string s, with all leading and trailing white space removed, as defined by Unicode.
//
// > (TrimSpaceは、Unicodeで定義されているように、すべての先頭と末尾の空白を除去した文字列sのスライスを返します。)
//
// # REFERENCES
//   - https://pkg.go.dev/strings@go1.22.1#TrimSpace
func TrimSpace() error {
	var (
		withSpace = "   hello world    "
		noSpace   = "hello world"
	)

	output.Stdoutf("[withSpace]", "%q\n", strings.TrimSpace(withSpace))
	output.Stdoutf("[noSpace  ]", "%q\n", strings.TrimSpace(noSpace))

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: string_trim_space

	   [Name] "string_trim_space"
	   [withSpace]          "hello world"
	   [noSpace  ]          "hello world"


	   [Elapsed] 24.94µs
	*/

}

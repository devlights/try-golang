package formatting

import (
	"fmt"

	"github.com/devlights/gomy/output"
)

// PaddingArbitaryLength は、文字列をパディングする際の桁数を外から指定するサンプルです.
//
// # REFERENCES
//   - https://golang.cafe/blog/golang-string-padding-example.html
//   - https://pkg.go.dev/fmt@latest
func PaddingArbitaryLength() error {
	// * (アスタリスク) を指定することで、外からパディング用桁数を指定出来る
	// アスタリスクの仕様については、 adverb_asterisk.go を参照
	var (
		s = "12345"
		i = 10
	)

	output.Stdoutl("[%*s]", fmt.Sprintf("'%*s'", i, s))

	return nil
}

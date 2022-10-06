package scanop

import (
	"fmt"

	"github.com/devlights/gomy/output"
)

// ReadFormattedInput は、fmt.Scanf() で書式化された入力値を読み取るサンプルです.
//
// # REFERENCES
//
//   - https://dev.to/azure/go-from-the-beginning-reading-user-input-i79
//   - https://pkg.go.dev/fmt@go1.19.2#Scanf
func ReadFormattedInput() error {
	var (
		value1 int
		value2 string
		format = "%3d %5s\n"
	)

	fmt.Print("INPUT(\\d{1,3} [^ ]{1,5}): ")

	// フォーマットに従った形で読み取りが行われる。
	// 改行もパターンの一つとして認識される。
	// フォーマットに合致しない場合 error となる。
	n, err := fmt.Scanf(format, &value1, &value2)
	if err != nil {
		return err
	}

	output.Stdoutf("[fmt.Scanf]", "count=%d\tvalue1=%d\tvalue2=%s\n", n, value1, value2)

	return nil
}

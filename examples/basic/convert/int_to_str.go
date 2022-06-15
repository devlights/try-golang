package convert

import (
	"fmt"

	"github.com/devlights/gomy/output"
)

// IntToStr は、fmt.Sprint() を使って、数値 (int) を 文字列 (string) に変換するサンプルです.
//
// REFERENCES:
//   - https://dave.cheney.net/2018/07/12/slices-from-the-ground-up
func IntToStr() error {
	var (
		i int     = 100
		f float32 = 12.345
	)

	s := fmt.Sprint(i)
	output.Stdoutf("[int to str]", "%[1]v(%[1]T) --> %[2]q(%[2]T)\n", i, s)

	s = fmt.Sprint(f)
	output.Stdoutf("[float to str]", "%[1]v(%[1]T) --> %[2]q(%[2]T)\n", f, s)

	return nil
}

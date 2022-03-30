package strconvs

import (
	"strconv"

	"github.com/devlights/gomy/output"
)

// BinToDecimal -- 2進数文字列を10進数に変換するサンプルです.
//
// REFERENCES
//  - https://pkg.go.dev/strconv
func BinToDecimal() error {
	var (
		value   = "10101010" // dec: 170
		parsed  int64
		intVal  int
		base    = 2
		bitSize = 0 // 0を指定すると int に収まるようにしてくれる
		err     error
	)

	parsed, err = strconv.ParseInt(value, base, bitSize)
	if err != nil {
		return err
	}

	intVal = int(parsed)

	output.Stdoutl("[bin]", value)
	output.Stdoutl("[dec]", intVal)

	return nil
}

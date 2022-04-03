package strconvs

import (
	"strconv"

	"github.com/devlights/gomy/output"
)

// DecToDec -- 10進数文字列を10進数に変換するサンプルです.
//
// REFERENCES
//  - https://pkg.go.dev/strconv
func DecToDec() error {
	var (
		values = []string{
			"255",
			"3735928559",
		}
	)

	var (
		parseInt = func(s string) int64 {
			v, _ := strconv.ParseInt(s, 10, 64)
			return v
		}
		atoi = func(s string) int64 {
			// strconv.Atoi() は strconv.ParseInt(v, 10, 0) と同じ
			v, _ := strconv.Atoi(s)
			return int64(v)
		}
	)

	for _, v := range values {
		output.Stdoutl("[original]", v)
		output.Stdoutl("[parseInt]", parseInt(v))
		output.Stdoutl("[atoi    ]", atoi(v))
		output.StderrHr()
	}
	return nil
}

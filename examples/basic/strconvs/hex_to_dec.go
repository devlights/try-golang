package strconvs

import (
	"strconv"

	"github.com/devlights/gomy/output"
)

// HexToDec -- 16進数文字列を10進数に変換するサンプルです.
func HexToDec() error {
	var (
		values = []string{
			"ff",
			"deadbeef",
		}
	)

	for _, v := range values {
		var (
			parsed int64
			err    error
		)

		// ParseInt() の 第２引数 base に 0 以外の値を指定している場合
		// prefix 付きの文字列を指定するとエラーとなる.
		// (ex: 0xff はエラーとなる。 ff はOK)
		parsed, err = strconv.ParseInt(v, 16, 64)
		if err != nil {
			return err
		}

		output.Stdoutl("[original]", v)
		output.Stdoutl("[parsed  ]", parsed)
		output.StdoutHr()
	}
	return nil
}

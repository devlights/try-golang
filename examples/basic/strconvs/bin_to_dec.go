package strconvs

import (
	"strconv"

	"github.com/devlights/gomy/output"
)

// BinToDec -- 2進数文字列を10進数に変換するサンプルです.
//
// REFERENCES
//   - https://pkg.go.dev/strconv
func BinToDec() error {
	var (
		values = []string{
			"11111111",
			"11011110101011011011111011101111",
		}
	)

	for _, v := range values {
		var (
			parsed int64
			err    error
		)

		// ParseInt() の 第２引数 base に 0 以外の値を指定している場合
		// prefix 付きの文字列を指定するとエラーとなる.
		// (ex: 0b1111 はエラーとなる。 1111 はOK)
		parsed, err = strconv.ParseInt(v, 2, 64)
		if err != nil {
			return err
		}

		output.Stdoutl("[original]", v)
		output.Stdoutl("[parsed  ]", parsed)
		output.StdoutHr()
	}

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: strconvs_bin_to_dec

	   [Name] "strconvs_bin_to_dec"
	   [original]           11111111
	   [parsed  ]           255
	   --------------------------------------------------
	   [original]           11011110101011011011111011101111
	   [parsed  ]           3735928559
	   --------------------------------------------------


	   [Elapsed] 56.51µs
	*/

}

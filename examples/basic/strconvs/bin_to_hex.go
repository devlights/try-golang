package strconvs

import (
	"strconv"

	"github.com/devlights/gomy/output"
)

// BinToHex -- 2進数文字列から16進数文字列へ変換するサンプルです。
func BinToHex() error {
	//
	// ２進数から１６進数へ変換する場合は以下の２段階で変換する.
	//   1) strconv.ParseInt() で int へ
	//   2) strconv.FormatInt() で string へ
	//

	var (
		values = []string{
			"0b11111111",
			"0b11011110101011011011111011101111",
		}
	)

	for _, v := range values {
		var (
			parsed int64
			err    error
		)

		parsed, err = strconv.ParseInt(v, 0, 0)
		if err != nil {
			return err
		}

		var (
			converted = strconv.FormatInt(parsed, 16)
		)

		output.Stdoutl("[original]", v)
		output.Stdoutl("[parsed  ]", parsed)
		output.Stdoutl("[conveted]", converted)
		output.StdoutHr()
	}

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: strconvs_bin_to_hex

	   [Name] "strconvs_bin_to_hex"
	   [original]           0b11111111
	   [parsed  ]           255
	   [conveted]           ff
	   --------------------------------------------------
	   [original]           0b11011110101011011011111011101111
	   [parsed  ]           3735928559
	   [conveted]           deadbeef
	   --------------------------------------------------


	   [Elapsed] 1.44571ms
	*/

}

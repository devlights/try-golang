package strconvs

import (
	"strconv"

	"github.com/devlights/gomy/output"
)

// HexToBin -- 16進数から2進数文字列へ変換するサンプルです。
func HexToBin() error {
	//
	// １６進数から２進数へ変換する場合は以下の２段階で変換する.
	//   1) strconv.ParseInt() で int へ
	//   2) strconv.FormatInt() で string へ
	//

	var (
		values = []string{
			"0xff",
			"0xDEADBEEF",
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
			converted = strconv.FormatInt(parsed, 2)
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

	   ENTER EXAMPLE NAME: strconvs_hex_to_bin

	   [Name] "strconvs_hex_to_bin"
	   [original]           0xff
	   [parsed  ]           255
	   [conveted]           11111111
	   --------------------------------------------------
	   [original]           0xDEADBEEF
	   [parsed  ]           3735928559
	   [conveted]           11011110101011011011111011101111
	   --------------------------------------------------


	   [Elapsed] 96.96µs
	*/

}

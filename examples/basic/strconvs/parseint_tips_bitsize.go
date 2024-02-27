package strconvs

import (
	"strconv"

	"github.com/devlights/gomy/output"
)

// ParseIntTipsBitSize -- strconv.ParseInt() の第３引数 bitSize を指定する際のTipsです。
//
// REFERENCES
//   - https://pkg.go.dev/strconv@latest#ParseInt
func ParseIntTipsBitSize() error {
	//
	// strconv.ParseInt() の 第３引数 bitSize には
	// 通常、値が収まるビットサイズを指定することになるが
	// 0を指定すると、 int に収まるようにしてくれる。
	//
	// 大抵の場合、intで取得したいときが多いので
	// 0を指定しておくと楽。
	//

	var (
		value   = "ff"
		base    = 16
		bitSize = 0
		parsed  int64
		err     error
	)

	parsed, err = strconv.ParseInt(value, base, bitSize)
	if err != nil {
		return err
	}

	output.Stdoutl("[original]", value)
	output.Stdoutl("[parsed  ]", int(parsed))

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: strconvs_parseint_tips_bitsize

	   [Name] "strconvs_parseint_tips_bitsize"
	   [original]           ff
	   [parsed  ]           255


	   [Elapsed] 21.09µs
	*/

}

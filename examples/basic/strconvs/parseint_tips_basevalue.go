package strconvs

import (
	"strconv"

	"github.com/devlights/gomy/output"
)

// ParseIntTipsBaseValue -- strconv.ParseInt() の第２引数 base を指定する際のTipsです。
//
// REFERENCES
//  - https://pkg.go.dev/strconv@latest#ParseInt
func ParseIntTipsBaseValue() error {
	//
	// strconv.ParseInt() の 第２引数 base には
	// 通常、元の文字列の進数を設定する。（2, 8, 10, 16 など)
	// この場合、0xff のように prefix を付けているとエラーとなる。
	//
	// しかし、base に 0 を指定した場合は prefix を付けていてもエラーにならない。
	// 2進数の場合は 0b 、8進数の場合は 0o 、16進数の場合は 0x となる。
	// この場合は逆に prefix が付いていないとエラーとなる。
	//
	// さらに、 parseint_tips_bitsize.go にあるように 第３引数の bitSize にも
	// 0 を指定することができる。
	//

	var (
		values = []string{
			"0b11111111",
			"0o377",
			"0xff",
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

		output.Stdoutl("[original]", v)
		output.Stdoutl("[parsed  ]", parsed)
		output.StdoutHr()
	}

	return nil
}

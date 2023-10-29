package floatop

import (
	"github.com/devlights/gomy/output"
	"github.com/shopspring/decimal"
)

// UsingDecimalPkg は、小数点計算を github.com/shopspring/decimal パッケージを利用して処理するサンプルです。
//
// # REFERENCES
//   - https://engineering.mercari.com/blog/entry/20201203-basis-point/
//   - https://github.com/shopspring/decimal
//
// # SEE ALSO
//   - examples/basic/floatop/rounding_error.go
func UsingDecimalPkg() error {
	var (
		v = decimal.RequireFromString("0")
	)

	for i := 0; i < 1000; i++ {
		v = v.Add(decimal.RequireFromString(".01"))
	}

	output.Stdoutl("[result]", v.StringFixed(1))

	return nil
}

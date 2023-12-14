package floatop

import (
	"math/big"

	"github.com/devlights/gomy/output"
)

// RoundingError は、小数点計算において近似値が利用され丸め誤差が出るサンプルです。
//
// # REFERENCES
//   - https://engineering.mercari.com/blog/entry/20201203-basis-point/
func RoundingError() error {
	//
	// 0.01は２進数で正確に表現できないため、近似値が利用されて丸め誤差が生じる
	//
	var (
		v float64
	)

	for i := 0; i < 1000; i++ {
		v += 0.01
	}

	output.Stdoutl("[float64]", v) // 9.999999999999831

	//
	// math/big を利用して同様の処理を行う
	//
	var (
		v2 = new(big.Rat)
	)

	for i := 0; i < 1000; i++ {
		v2 = new(big.Rat).Add(v2, big.NewRat(1, 100))
	}

	output.Stdoutl("[big.Rat]", v2.FloatString(1)) // 10.0

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: floatop_rounding_error

	   [Name] "floatop_rounding_error"
	   [float64]            9.999999999999831
	   [big.Rat]            10.0


	   [Elapsed] 1.32107ms
	*/

}

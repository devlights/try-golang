package floatop

import (
	"math/rand"

	"github.com/devlights/gomy/output"
)

// OrderOfComputation -- 浮動小数点は計算の順序によって結果が変わることのサンプルです.
//
// # REFERENCES
//   - https://zenn.dev/kumackey/articles/d20230708-a7c195db087338
//   - https://wp.jmuk.org/2023/06/21/%e6%b5%ae%e5%8b%95%e5%b0%8f%e6%95%b0%e7%82%b9%e6%95%b0%e3%81%ae%e5%8a%a0%e7%ae%97%e3%81%ae%e9%a0%86%e5%ba%8f%e3%81%ab%e3%83%8f%e3%83%9e%e3%81%a3%e3%81%9f%e8%a9%b1/
//   - https://www.jpcert.or.jp/sc-rules/c-flp01-c.html
func OrderOfComputation() error {
	//
	// 浮動小数点は計算の順序によって結果が変わる
	// ただし、float32の場合は計算結果をfloat64で受ければ大丈夫
	//
	var (
		m = make(map[int]float32)
	)

	for i := 0; i < 1000; i++ {
		v := rand.Float32()
		m[i] = v
	}

	for i := 0; i < 10; i++ {
		var (
			total32 float32
			total64 float64
		)

		for _, v := range m {
			total32 += v
			total64 += float64(v)
		}

		output.Stdoutf("[Total (float32, float64)]", "%v\t%v\n", total32, total64)
	}

	return nil

	/*
	   $ task
	   task: Task "build" is up to date
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: floatop_order_of_computation

	   [Name] "floatop_order_of_computation"
	   [Total (float32, float64)] 501.02402    501.02436977016623
	   [Total (float32, float64)] 501.02417    501.02436977016623
	   [Total (float32, float64)] 501.0243     501.02436977016623
	   [Total (float32, float64)] 501.02444    501.02436977016623
	   [Total (float32, float64)] 501.02475    501.02436977016623
	   [Total (float32, float64)] 501.02423    501.02436977016623
	   [Total (float32, float64)] 501.02454    501.02436977016623
	   [Total (float32, float64)] 501.02466    501.02436977016623
	   [Total (float32, float64)] 501.02435    501.02436977016623
	   [Total (float32, float64)] 501.02426    501.02436977016623


	   [Elapsed] 586.92µs
	*/

}

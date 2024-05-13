package zerovalues

import "github.com/devlights/gomy/output"

// Array は、Goにおける 配列 のゼロ値についてのサンプルです.
//
// # REFERENCES
//   - https://go.dev/tour/basics/12
//   - https://brain2life.hashnode.dev/default-zero-values-in-go
func Array() error {
	//
	// 配列 の ゼロ値 は、その配列の基底型のゼロ値が指定要素数分設定されている状態.
	//
	var (
		a [10]int
	)

	output.Stdoutf("[array zerovalue	]", "%v\n", a)

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: zerovalues_array

	   [Name] "zerovalues_array"
	   [array zerovalue        ]   [0 0 0 0 0 0 0 0 0 0]


	   [Elapsed] 12.31µs
	*/

}

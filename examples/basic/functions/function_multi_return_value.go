package functions

import "fmt"

// Goの関数は戻り値を複数定義することが可能
// 複数の戻り値が存在する場合は、戻り値をカッコで囲む
func toDoubleTheNumber(x int) (int, int) {
	return x, x * 2
}

// FunctionMultiReturnValue -- 複数の戻り値を持つ関数を定義できることを確認するサンプルです。
func FunctionMultiReturnValue() error {

	var (
		x = 100
	)

	original, result := toDoubleTheNumber(x)

	fmt.Printf("toDoubleTheNumber(100) ==> (%d, %d)\n", original, result)
	return nil

	/*
	   $ task
	   task: Task "build" is up to date
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: function_multi_return_value

	   [Name] "function_multi_return_value"
	   toDoubleTheNumber(100) ==> (100, 200)


	   [Elapsed] 32.74µs
	*/

}

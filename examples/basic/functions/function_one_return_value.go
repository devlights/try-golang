package functions

import "fmt"

// Goの関数定義は戻り値を後に配置するスタイル
// 戻り値が一つの場合は、カッコで囲む必要はない
func add(x, y int) int {
	return x + y
}

// OneReturnValue -- 戻り値が一つの関数定義のサンプル
func OneReturnValue() error {

	var (
		x, y = 10, 20
	)

	fmt.Printf("add(10,20) ==> %d\n", add(x, y))

	return nil

	/*
	   $ task
	   task: Task "build" is up to date
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: function_one_return_value

	   [Name] "function_one_return_value"
	   add(10,20) ==> 30


	   [Elapsed] 13.72µs
	*/

}

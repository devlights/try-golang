package functions

import "fmt"

// Goの関数定義は戻り値を後に配置するスタイル
// 戻り値が一つの場合は、カッコで囲む必要はない
func add(x, y int) int {
	return x + y
}

// FunctionOneReturnValue -- 戻り値が一つの関数定義のサンプル
func FunctionOneReturnValue() error {

	var (
		x, y = 10, 20
	)

	fmt.Printf("add(10,20) ==> %d\n", add(x, y))

	return nil
}

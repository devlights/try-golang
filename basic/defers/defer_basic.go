package defers

import "fmt"

// Basic - defer の基本的な使い方についてのサンプルです。
//
// deferに匿名関数を作って処理する場合はクロージャで渡す必要がある。
// クロージャは自身が内部で利用してる外側の変数の値をバインドしているので
// 最終的に呼び出される時点での値が利用される。
// deferに関数呼び出しを指定する場合は呼び出した時点での値が渡ることになる。
func Basic() error {
	i := 1

	defer func() {
		fmt.Printf("defer1 -- %d\n", i)
	}()

	defer printVal(i)

	i++

	defer func() {
		fmt.Printf("defer2 -- %d\n", i)
	}()

	defer printVal(i)

	i++

	return nil

	// Output:
	// printVal() -- 2
	// defer2 -- 3
	// printVal() -- 1
	// defer() -- 3
}

func printVal(val int) {
	fmt.Printf("printVal() -- %d\n", val)
}

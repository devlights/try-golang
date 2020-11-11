package effectivego12

import "fmt"

// Arrays -- Effective Go - Arrays の 内容についてのサンプルです。
func Arrays() error {
	/*
		https://golang.org/doc/effective_go.html#arrays

		- 配列は値型。なので、別の配列に代入すると値のコピーが行われる。
		- 配列をそのまま関数の引数として渡すと値がコピーされる。関数内で値を変更する場合はポインタで渡す。
		- Goでは配列の要素数は配列の型の一部。なので、 [2]int と [3]int は全く別の型となる。
	*/
	var (
		dump = func(arr [3]int, message string) {
			fmt.Printf("[%-25s] arr:%v\n", message, arr)
		}

		dumpP = func(arr *[3]int, message string) {
			fmt.Printf("[%-25s] arr:%v\n", message, *arr)
		}

		byVal = func(arr [3]int) {
			for i := 0; i < len(arr); i++ {
				arr[i] = arr[i] + 10
			}

			dump(arr, "byVal")
		}

		byRef = func(arr *[3]int) {
			for i := 0; i < len(arr); i++ {
				arr[i] = arr[i] + 10
			}

			dumpP(arr, "byRef")
		}
	)

	var (
		arr1 = [3]int{1, 2, 3}
	)

	// Goでは配列数は型の一部。[3]int と [2]int は異なる型
	// arr1 = [2]int{}

	dump(arr1, "init")

	// 値渡し
	byVal(arr1)
	dump(arr1, "after byVal()")

	// 参照渡し
	byRef(&arr1)
	dump(arr1, "after byRef()")

	return nil
}

package loops

import "fmt"

// RangeLoopTmpValueWithArray は、for range ループにて配列をループした際の注意するべき挙動についてのサンプルです.
func RangeLoopTmpValueWithArray() error {
	var (
		ary = [3]int{1, 2, 3}
	)
	defer func() { fmt.Printf("%v\n", ary) }()

	// rangeループでは、ループが始まる前に ary の値が一時変数に確保され
	// その一時変数を用いてループされる。なので、ループ内で元の ary の要素値を
	// 更新しても、ループ変数として渡される v の値は変わらない。
	// (変更される前に一時変数として確保されたものの値であるため)
	for _, v := range ary {
		ary[2] = 99
		fmt.Println(v)
	}

	return nil
}

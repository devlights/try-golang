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

	/*
	   $ task
	   task: Task "build" is up to date
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: loops_range_loop_tmpvalue_with_array

	   [Name] "loops_range_loop_tmpvalue_with_array"
	   1
	   2
	   3
	   [1 2 99]


	   [Elapsed] 40.151µs
	*/

}

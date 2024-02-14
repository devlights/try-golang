package slices

import "fmt"

// Basic02 -- スライスについてのサンプル
func Basic02() error {
	// スライスは make() からでも作れる
	// make() は、 slice, map, chan を生成する関数
	l1 := make([]int, 0, 5)
	fmt.Printf("l1: %T\n", l1)

	for i := 0; i < 5; i++ {
		l1 = append(l1, i)
	}

	for _, item := range l1 {
		fmt.Println(item)
	}

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: slice_basic02

	   [Name] "slice_basic02"
	   l1: []int
	   0
	   1
	   2
	   3
	   4


	   [Elapsed] 21.78µs
	*/

}

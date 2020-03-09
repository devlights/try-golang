package slices

import "fmt"

// スライスについてのサンプル
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
}

package slices

import "fmt"

// Basic05 -- 値渡しと参照渡しについて
func Basic05() error {
	// GO言語では 配列は値, スライスは参照となる
	// 配列をそのまま渡すと「値渡し」となる
	// スライスは参照なので同じ「値渡し」だが
	// 参照がコピーされて渡される（他の言語の配列と同じ）
	ary01 := [5]int{1, 2, 3, 4, 5}
	sli01 := []int{1, 2, 3, 4, 5}

	fmt.Printf("array: %v\n", ary01)
	fmt.Printf("slice: %v\n", sli01)

	calcArray(ary01) // 配列を値渡ししているので値は変化しない
	calcSlice(sli01) // スライスなので値が変化する

	fmt.Printf("array: %v\n", ary01)
	fmt.Printf("slice: %v\n", sli01)

	calcArrayPointer(&ary01) // ポインタで渡すと、当然値は変化する
	fmt.Printf("array: %v\n", ary01)

	return nil
}

func calcArray(val [5]int) {
	for i := 0; i < len(val); i++ {
		val[i] = val[i] * 2
	}
}

func calcArrayPointer(val *[5]int) {
	for i := 0; i < len(val); i++ {
		val[i] = val[i] * 2
	}
}

func calcSlice(val []int) {
	for i := 0; i < len(val); i++ {
		val[i] = val[i] * 2
	}
}

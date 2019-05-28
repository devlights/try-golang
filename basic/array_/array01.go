package array_

import "fmt"

func Array01() error {
	// サイズが固定化されているものが配列
	// スライスは配列の一部分を指し示す参照
	var array1 [3]int

	// 配列の要素数は len() で取得できる
	fmt.Println(len(array1))

	// 値はデフォルト値で初期設定される (この配列は int なので 0　で埋まる)
	printArray(&array1)

	// 値の設定
	array1[1] = 100
	printArray(&array1)

	// 初期化と同時に宣言して設定することも可能
	array2 := [...]int{1, 2, 3}
	printArray(&array2)

	// そのまま渡すと値のコピーが発生する
	updateArrayVal(array2) // => 値のコピーが発生するので、値が変わらない
	printArray(&array2)

	updateArrayPtr(&array2) // => ポインタ経由させるので値が変わる
	printArray(&array2)

	// スライスは配列の一部分を指し示す参照
	// 全体を指し示すスライスは以下のようにして取得できる
	slice1 := array2[:]

	// スライスを型情報出力してみると []int のように []の中に要素数が表示されない
	fmt.Printf("array: %T, slice: %T\n", array2, slice1)

	updateSliceVal(slice1) // => 参照なのでそのまま渡しても値が変わる (C# や Java などのリストと同じ)
	printSlice(&slice1)
	printArray(&array2) // => 元の配列も当然変わる

	array3 := [...]int{1, 2, 4}
	slice2 := array3[:]
	updateSlicePtr(&slice2) // => ポインタで渡しても同じことになる。ただし関数の中でポインタから実体に戻す必要がある
	printSlice(&slice2)
	printArray(&array3) // => 元の配列も当然変わる

	return nil
}

func printArray(ary *[3]int) {
	for _, item := range ary {
		fmt.Print(item, " ")
	}

	fmt.Println("")
}

func printSlice(slice *[]int) {
	for _, item := range *slice {
		fmt.Print(item, " ")
	}

	fmt.Println("")
}

func updateArrayVal(ary [3]int) {
	ary[1] = 999
}

func updateArrayPtr(ary *[3]int) {
	ary[1] = 999
}

func updateSliceVal(slice []int) {
	slice[1] = 998
}

func updateSlicePtr(slice *[]int) {
	(*slice)[1] = 997
}

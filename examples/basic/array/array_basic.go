package array

import (
	"strconv"
	"strings"

	"github.com/devlights/gomy/output"
)

// Basic は、Goにおける配列の基本的な使い方についてのサンプルです.
//
// REFERENCES:
//   - https://dave.cheney.net/2018/07/12/slices-from-the-ground-up
func Basic() error {
	// サイズが固定化されているものが配列
	// スライスは配列の一部分を指し示す参照
	var array1 [3]int

	// 配列の要素数は len() で取得できる
	output.Stdoutl("len(array1)", len(array1))

	// 値はデフォルト値で初期設定される (この配列は int なので 0　で埋まる)
	printArray("ary (1)", &array1)

	// 値の設定
	array1[1] = 100
	printArray("ary (2)", &array1)

	// 初期化と同時に宣言して設定することも可能
	array2 := [...]int{1, 2, 3}
	printArray("ary (3)", &array2)

	// そのまま渡すと値のコピーが発生する
	updateArrayVal(array2) // => 値のコピーが発生するので、値が変わらない
	printArray("ary (4)", &array2)

	updateArrayPtr(&array2) // => ポインタ経由させるので値が変わる
	printArray("ary (5)", &array2)

	// スライスは配列の一部分を指し示す参照
	// 全体を指し示すスライスは以下のようにして取得できる
	slice1 := array2[:]

	// スライスを型情報出力してみると []int のように []の中に要素数が表示されない
	output.Stdoutf("type", "array: %T, slice: %T\n", array2, slice1)

	updateSliceVal(slice1) // => 参照なのでそのまま渡しても値が変わる (C# や Java などのリストと同じ)
	printSlice("sli (1)", &slice1)
	printArray("ary (6)", &array2) // => 元の配列も当然変わる

	array3 := [...]int{1, 2, 4}
	slice2 := array3[:]
	updateSlicePtr(&slice2) // => ポインタで渡しても同じことになる。ただし関数の中でポインタから実体に戻す必要がある
	printSlice("sli (2)", &slice2)
	printArray("ary (7)", &array3) // => 元の配列も当然変わる

	return nil

	/*
	   $ task
	   task: Task "build" is up to date
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: array_basic_usage

	   [Name] "array_basic_usage"
	   len(array1)          3
	   ary (1)              0 0 0
	   ary (2)              0 100 0
	   ary (3)              1 2 3
	   ary (4)              1 2 3
	   ary (5)              1 999 3
	   type                 array: [3]int, slice: []int
	   sli (1)              1 998 3
	   ary (6)              1 998 3
	   sli (2)              1 997 4
	   ary (7)              1 997 4


	   [Elapsed] 97.009µs
	*/

}

func printArray(prefix string, ary *[3]int) {

	sb := strings.Builder{}
	for _, item := range ary {
		s := strconv.Itoa(item)
		_, _ = sb.WriteString(s)
		_, _ = sb.WriteString(" ")
	}

	output.Stdoutl(prefix, sb.String())
}

func printSlice(prefix string, slice *[]int) {

	sb := strings.Builder{}
	for _, item := range *slice {
		s := strconv.Itoa(item)
		_, _ = sb.WriteString(s)
		_, _ = sb.WriteString(" ")
	}

	output.Stdoutl(prefix, sb.String())
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

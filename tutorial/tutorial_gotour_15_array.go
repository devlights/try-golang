package tutorial

import "fmt"

// Array は、 Tour of Go - Arrays (https://tour.golang.org/moretypes/6) の サンプルです。
func Array() error {
	// ------------------------------------------------------------
	// Go言語の配列
	// Go言語の配列は、 [要素数]型名 という形で宣言する.
	// 特殊なのが、配列の長さ（要素数）は「型の一部」であるということ。
	// なので、 例えば 同じ int の配列であっても
	//   - [2]int
	//   - [3]int
	// は異なる型となる。これは不便だと最初思えるが、Go言語にはスライスという
	// 概念があるので、実際には気にならない.
	//
	// スライスは「可変長」で、配列は「固定長」である.
	// ------------------------------------------------------------
	var (
		arr [2]int
	)

	// 初期値は設定していない場合、その要素型のゼロ値となる
	fmt.Println(arr)

	// 配列の長さは 組み込み関数 len で取れる
	fmt.Println(len(arr))

	// Cの様に配列をループする場合は以下のようにする
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%v\t", arr[i])
	}

	fmt.Println("")

	// foreach ループする場合は以下のようにする
	for _, v := range arr {
		fmt.Printf("%v\t", v)
	}

	fmt.Println("")

	// 値の設定も他の言語と同じ
	arr[0] = 100
	arr[1] = 200
	for _, v := range arr {
		fmt.Printf("%v\t", v)
	}

	fmt.Println("")

	return nil
}

package gotour15

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
	//
	// C言語経験者だと、引っかかってしまう点が一つあり
	// Goの配列は、値であるので関数に引数として配列を渡した場合
	// 関数内で引数の配列の要素に対して値更新をしても、呼び出し元の
	// 配列には影響しないという点がある。同じようにするには配列のポインタを
	// 渡すようにする。
	//
	// C#などの経験者も、配列を関数の引数で渡す場合、頭では参照が渡っていると
	// 認識してしまうので注意が必要。同じような動きをさせる場合はポインタで渡すこと。
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

	// Goの配列は値なので、そのまま渡すと配列自体がコピーされて渡る。
	// そのため、関数内で配列の値を編集しても呼び元には影響しない。
	fmt.Printf("[updateArray前] %v\n", arr)
	updateArray(arr)
	fmt.Printf("[updateArray後] %v\n", arr)

	fmt.Println("")

	// ポインタで渡すと望んだ動きとなる
	fmt.Printf("[updateArray2前] %v\n", arr)
	updateArray2(&arr)
	fmt.Printf("[updateArray2後] %v\n", arr)

	return nil
}

func updateArray2(ints *[2]int) {
	// 本来、ポインタでパラメータが渡されているので
	// 値にアクセスするには、まずデリファレンスが必要となる
	// (*ints)[0] = 999
	// しかし、Goのランタイム側がこれを吸収してくれるので
	// 普通にデリファレンス無しで配列操作のように書くことができる
	ints[0] = 999
	ints[1] = 998

	fmt.Printf("[updateArray2中] %v\n", ints)
}

func updateArray(ints [2]int) {
	ints[0] = 999
	ints[1] = 998

	fmt.Printf("[updateArray中] %v\n", ints)
}

package tutorial

import "fmt"

// Pointer は、 Tour of Go - Pointers (https://tour.golang.org/moretypes/1) の サンプルです。
func Pointer() error {
	// ------------------------------------------------------------
	// Go言語のポインタ
	//
	// Go言語は、C や C++ と同様に ポインタ型 を持っている
	// 使い方もほぼ同じであるが、ポインタ演算は出来ないようになっている
	// ------------------------------------------------------------
	//noinspection GoVarAndConstTypeMayBeOmitted
	var (
		i   int         = 100
		p   *int        = &i
		arr [2]int      = [2]int{0, 0} // Go言語では、配列は要素数も含めて型となる。[1]int　と [2]int は同じ型ではない.
		sli []int       = []int{0, 0}
		m   map[int]int = map[int]int{1: 0, 2: 0}
	)

	showValue(i, p)

	updateValue(p, 200)
	showValue(i, p)

	updateValue(&i, 300)
	showValue(i, p)

	// 配列
	fmt.Println(arr)
	updateArrayNotPointer(arr)
	fmt.Println(arr)
	updateArrayPointer(&arr)
	fmt.Println(arr)

	// スライス
	fmt.Println(sli)
	updateSliceNotPointer(sli)
	fmt.Println(sli)
	updateSlicePointer(&sli)
	fmt.Println(sli)

	// マップ
	fmt.Println(m)
	updateMapNotPointer(m)
	fmt.Println(m)
	updateMapPointer(&m)
	fmt.Println(m)

	return nil
}

func showValue(i int, p *int) {
	fmt.Printf("i=%d\tp=%p\t*p=%d\n", i, p, *p)
}

func updateValue(p *int, v int) {
	*p = v
}

func updateArrayNotPointer(arr [2]int) {
	// ポインタで渡していないので、値がコピーされて関数に渡る
	// なので、値を変更しても元の配列には影響しない
	arr[0] = 100
	arr[1] = 200
}

func updateArrayPointer(arr *[2]int) {
	// ポインタで渡しているので、参照（ポインタ）がコピーされた関数に渡る
	// なので、値を変更すると元の配列にも影響する
	// 本来であれば、arrは配列へのポインタなのでインデックス操作が出来ない
	// はずであるが、配列の場合、言語側で特殊措置が行われるため
	// ポインタであっても、そのままインデックス操作が行えるようになっている
	arr[0] = 100
	arr[1] = 200
}

func updateSliceNotPointer(sli []int) {
	// Go言語のスライスは、内部に参照先へのポインタを持っている構造になっている
	// なので、ポインタを取得して渡さなくても、値のコピー時に内部のポインタがコピーされるため
	// 結果として、同じ参照先を更新することになる。
	sli[0] = 100
	sli[1] = 200
}

func updateSlicePointer(sli *[]int) {
	// わざわざスライスのポインタを取得すると、一旦 ポインタの実態 を取り出す手間が必要となる
	// スライスは、内部に参照先へのポインタを持っているため、ポインタ取得せずそのまま関数に渡しても問題ない
	(*sli)[0] = 300
	(*sli)[1] = 400
}

func updateMapNotPointer(m map[int]int) {
	// マップもスライスと同様に内部に参照先へのポインタを持っているので
	// わざわざポインタを取得して渡す必要はない
	m[1] = 100
	m[2] = 200
}

func updateMapPointer(m *map[int]int) {
	// スライスと同様に、マップのポインタからは直接 マップ のキー指定（インデックス指定）
	// が出来ないので、一旦デリファレンスをする必要がある
	(*m)[1] = 300
	(*m)[2] = 400
}

package slices

import "fmt"

// Pointer は、スライスの ポインタ 利用時についてのサンプルです.
//
// REFERENCES:
//   - https://dave.cheney.net/2018/07/12/slices-from-the-ground-up
func Pointer() error {
	// ----------------------------------------------------------------
	// スライスのポインタ利用について
	//
	// スライスは内部構造で参照先となっている配列のポインタを持っているため
	// 通常ポインタを利用する必要はないと記載されている場合があるが
	// C#やJavaなどでよくやるように、メソッドにリストを渡して、メソッド内部で
	// そのリストに対して、要素を追加したりする場合、スライスをポインタで渡して
	// 処理したほうが良い場合がある。
	//
	// スライスは以下のような構造となっている。
	// (https://golang.org/reflect/#SliceHeader)
	//
	// type SliceHeader struct {
	//     Data uintptr
	//     Len  int
	//     Cap  int
	// }
	//
	// 関数に、スライスをそのまま渡すと、この構造体が値コピーされて渡る。
	// 内部にデータへのポインタを持っているので、そのまま渡しても参照先の配列
	// は同じポインタを示すので、通常は問題ない。
	//
	// しかし、渡された関数側でスライスの要素を追加したりする場合は
	// 元々、スライス作成時に指定した cap を超える場合がある。
	// その場合、Goは自動で新しい配列を用意してDataのポインタを変更してくれる。
	// この挙動により、関数内部でスライスをいじった場合に呼び出し元と関数側の
	// Dataフィールドのポインタが異なってしまうため、追加された要素が反映されない。
	//
	// また、cap を超えない場合でも、スライス自身を値渡ししているため
	// Capフィールドが更新されないので、呼び元で値を見ても反映されていないように
	// 見える。この場合は以下のようにすることで、最新のcapを適用することができる。
	//
	//     sli := sli[:cap(sli)]
	//
	// 上記の問題は、どちらも スライス を ポインタ で渡すことで解決はできる。
	// ----------------------------------------------------------------
	// スライスを値渡しで受け取って、要素を追加
	dump := func(sli []int, prefix string) {
		fmt.Printf("[%-25s] len:%d\tcap:%d\tvalues:%v\n", prefix, len(sli), cap(sli), sli)
	}

	byVal := func(sli []int, count int) {
		for i := 0; i < count; i++ {
			sli = append(sli, i)
		}

		dump(sli, "byVal call")
	}

	byRef := func(sli *[]int, count int) {
		for i := 0; i < count; i++ {
			*sli = append(*sli, i)
		}

		dump(*sli, "byRef call")
	}

	// ---------------------------------------
	// スライスを値渡しの場合
	//   int の スライスを作成（ cap は 1 )
	// ---------------------------------------
	sliByVal := make([]int, 0, 1)
	dump(sliByVal, "sliByVal init")

	// 要素を一つ追加
	//   内部で要素が追加されるがcapが更新されていないため
	//   呼び元でそのまま見ると、追加された要素が見えない。
	//   capに到達していないので、capを更新すると追加された要素が見える.
	byVal(sliByVal, 1)
	dump(sliByVal, "sliByVal - byVal(1) - 1")

	// cap を更新
	sliByVal = sliByVal[:cap(sliByVal)]
	dump(sliByVal, "sliByVal - update cap")

	// 要素をさらに一つ追加
	//   内部で要素が追加されるタイミングでcapに到達するため
	//   goは自動的に新しい配列を用意してスライス内部の配列の
	//   参照を書き換える. 呼び元のスライスが元々参照していた
	//   データ配列の参照とは異なる状態になっているため
	//   あとで、capを更新しても追加された要素は見ることができない。
	//   （別の配列に変わってしまっているため）
	byVal(sliByVal, 1)
	dump(sliByVal, "sliByVal - byVal(1) - 2")

	// cap を更新
	sliByVal = sliByVal[:cap(sliByVal)]
	dump(sliByVal, "sliByVal - update cap")

	fmt.Println("-----------------------------")

	// ---------------------------------------
	// スライスをポインタ渡しの場合
	//   int の スライスを作成 ( cap は 1 )
	// ---------------------------------------
	sliByRef := make([]int, 0, 1)
	dump(sliByRef, "sliByRef init")

	// 要素を一つ追加
	//   内部で要素が追加されるが、スライス自身をポインタで
	//   渡しているため、capの更新もそのまま見えている。
	//   なので、 cap の更新は必要なく、追加された要素も見える。
	byRef(&sliByRef, 1)
	dump(sliByRef, "sliByRef - byRef(1) - 1")

	// 要素をさらに一つ追加
	//   内部で要素が追加されるタイミングでcapに到達するため
	//   goは自動的に新しい配列を用意してスライス内部の配列の
	//   参照を書き換える. 呼び元のスライスが元々参照していた
	//   データ配列の参照とは異なる状態になっているが
	//   スライス自身をポインタで渡しているため、その変更は
	//   呼び元のスライスにもそのまま適用されている。
	//   なので、データは何もせずとも更新後がちゃんと見える。
	byRef(&sliByRef, 1)
	dump(sliByRef, "sliByRef - byRef(1) - 2")

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: slice_pointer

	   [Name] "slice_pointer"
	   [sliByVal init            ] len:0       cap:1   values:[]
	   [byVal call               ] len:1       cap:1   values:[0]
	   [sliByVal - byVal(1) - 1  ] len:0       cap:1   values:[]
	   [sliByVal - update cap    ] len:1       cap:1   values:[0]
	   [byVal call               ] len:2       cap:2   values:[0 0]
	   [sliByVal - byVal(1) - 2  ] len:1       cap:1   values:[0]
	   [sliByVal - update cap    ] len:1       cap:1   values:[0]
	   -----------------------------
	   [sliByRef init            ] len:0       cap:1   values:[]
	   [byRef call               ] len:1       cap:1   values:[0]
	   [sliByRef - byRef(1) - 1  ] len:1       cap:1   values:[0]
	   [byRef call               ] len:2       cap:2   values:[0 0]
	   [sliByRef - byRef(1) - 2  ] len:2       cap:2   values:[0 0]


	   [Elapsed] 138.06µs
	*/

}

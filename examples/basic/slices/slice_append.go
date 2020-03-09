package slices

import "fmt"

// SliceAppend は、スライスの append 利用時についてのサンプルです.
func SliceAppend() error {
	// ----------------------------------------------------------------
	// Go の スライス に対しての append の利用
	//
	// スライスに append 関数を使って要素を追加していく際に
	// 以下の注意点がある。
	//
	// - append 関数は、追加を行う際に元の配列に要素が足りない場合は要素を増やした
	//   別の配列を用意して、そちらに移し替えたものを戻り値として返す.
	//
	// なので、よく別の言語でやるように、関数の引数にリストを渡して関数内の処理で
	// そのリストに要素を追加していくような形の処理をGoでそのまま書くと、元のリストには
	// 何の要素も追加されないという事態が発生するので注意が必要。
	//
	// 回避としては、元々のサイズをしっかり確保したスライスを作っておくか
	// 関数に渡すのではなく、関数内でスライスを作って戻り値で返して、呼び元で統合するなどがある.
	//
	// その他にも、関数内で引数で指定されたスライスに対して append すると、その関数内では
	// 値は更新されているが、呼び出し元で見てみると、更新されていない現象というのもある。
	// これは、同じ配列を示しているが、呼び出し元の slice の len が更新されていないから。
	// (https://christina04.hatenablog.com/entry/2017/09/26/190000)
	//
	// 呼び出し元のlenを更新するには
	//   ints = ints[:cap(ints)]
	// という呼び出しを入れる。これで、このスライスのlenが更新されることになる。
	//
	// 以下の例では、capがゼロなスライスと余裕も持っているスライスの2つを定義して
	// 関数の引数として渡して、その中で10個の要素をappendしている。
	// capがゼロのスライスは、その際に必ず拡張が発生するので、別のアドレスを指し示すように
	// 変わってしまう。
	// ----------------------------------------------------------------
	var (
		zeroCapSlice = make([]int, 0, 0)
		manyCapSlice = make([]int, 0, 15)
	)

	fmt.Printf("[zeroCap] len:%d\tcap:%d\taddr:[%p]\n", len(zeroCapSlice), cap(zeroCapSlice), zeroCapSlice)
	appendItems(zeroCapSlice)

	fmt.Printf("[manyCap] len:%d\tcap:%d\taddr:[%p]\n", len(manyCapSlice), cap(manyCapSlice), manyCapSlice)
	appendItems(manyCapSlice)

	// 別の関数内で要素を追加していても、呼び出し元のスライスから見るとlenが変わっていないため
	// 値が表示されない。以下のように 再度スライスをすることによって len が更新される。
	// (https://christina04.hatenablog.com/entry/2017/09/26/190000)
	zeroCapSlice = zeroCapSlice[:cap(zeroCapSlice)]
	manyCapSlice = manyCapSlice[:cap(manyCapSlice)]

	fmt.Println("----------------------------------------------")
	fmt.Printf("[zeroCap] len:%d\tcap:%d\taddr:[%p]\n", len(zeroCapSlice), cap(zeroCapSlice), zeroCapSlice)
	fmt.Printf("[manyCap] len:%d\tcap:%d\taddr:[%p]\n", len(manyCapSlice), cap(manyCapSlice), manyCapSlice)
	fmt.Println("zeroCapSlice", zeroCapSlice)
	fmt.Println("manyCapSlice", manyCapSlice)
	fmt.Println("----------------------------------------------")

	// 安全なのが、関数に引数で渡して更新するのではなく、関数内で別途スライスを作って戻り値として返すようにすること.
	// 呼び元は、受け取ったデータを自分で管理しているスライスに追加する.
	var (
		zeroCapSlice2 = make([]int, 0, 0)
		manyCapSlice2 = make([]int, 0, 15)
	)

	fmt.Printf("[zeroCap2] len:%d\tcap:%d\taddr:[%p]\n", len(zeroCapSlice2), cap(zeroCapSlice2), zeroCapSlice2)
	fmt.Printf("[manyCap2] len:%d\tcap:%d\taddr:[%p]\n", len(manyCapSlice2), cap(manyCapSlice2), manyCapSlice2)
	fmt.Println("zeroCapSlice2", zeroCapSlice2)
	fmt.Println("manyCapSlice2", manyCapSlice2)

	items := retrieveItems(10)
	for _, v := range items {
		zeroCapSlice2 = append(zeroCapSlice2, v)
		manyCapSlice2 = append(manyCapSlice2, v)
	}

	fmt.Printf("[zeroCap2] len:%d\tcap:%d\taddr:[%p]\n", len(zeroCapSlice2), cap(zeroCapSlice2), zeroCapSlice2)
	fmt.Printf("[manyCap2] len:%d\tcap:%d\taddr:[%p]\n", len(manyCapSlice2), cap(manyCapSlice2), manyCapSlice2)
	fmt.Println("zeroCapSlice2", zeroCapSlice2)
	fmt.Println("manyCapSlice2", manyCapSlice2)
	fmt.Println("----------------------------------------------")

	return nil
}

func retrieveItems(count int) []int {
	r := make([]int, 0, count)
	for i := 0; i < count; i++ {
		r = append(r, i)
	}

	return r
}

func appendItems(ints []int) {
	fmt.Printf("\t[append][before] len:%d\tcap:%d\taddr:[%p]\n", len(ints), cap(ints), ints)

	for i := 0; i < 10; i++ {
		ints = append(ints, i)
	}

	fmt.Printf("\t[append][after]  len:%d\tcap:%d\taddr:[%p]\n", len(ints), cap(ints), ints)
}

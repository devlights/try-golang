package formatting

import "fmt"

// AdverbExplicitArgumentIndexes -- フォーマッティングの Explicit argument indexes についてのサンプルです。
//
// REFERENCES:
//   - https://golang.org/fmt/
func AdverbExplicitArgumentIndexes() error {
	// ------------------------------------------------------------
	// Go のフォーマット仕様には
	//   Explicit argument indexes
	// という仕様がある。
	//
	// フォーマット文字列内に以下が指定できる
	//   - [1]とすると、添字内のパラメータの値を指定できる
	//     中の数字は、フォーマット文字列の後のパラメータの位置
	//     (最初のパラメータが 1 )
	//   - [1]* のようにアスタリスクをつけると、そのパラメータの
	//     値を指示子の値として利用できる
	// ------------------------------------------------------------

	// 以下の2つは同じ結果となる
	i := 10
	fmt.Printf("%d(%T)\n", i, i)
	fmt.Printf("%d(%[1]T)\n", i)

	// 以下の2つは同じ結果となる
	f := 32.8
	fmt.Printf("%3.2f\n", f)
	fmt.Printf("%[3]*.[2]*[1]f\n", f, 2, 3)

	// 複雑になるので、あまり意味はないが、以下の様にも出来る
	var (
		numberLength = 10
		totalLength  = numberLength + 5
	)

	fmt.Printf("% [2]*[1]s\n", fmt.Sprintf("%0[2]*[1]d\n", i, numberLength), totalLength)

	return nil
}

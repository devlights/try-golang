package array

import (
	"github.com/devlights/gomy/output"
)

// CopyFromSlice -- スライスから配列へコピーするサンプルです。
func CopyFromSlice() error {
	// ----------------------------------------------
	// スライスから配列へコピー
	//
	// ビルドインのcopy()を利用すれば良い。
	//
	// [参考]
	// https://stackoverflow.com/a/30285971
	// https://blog.golang.org/slices-intro
	// https://blog.golang.org/slices
	// ----------------------------------------------
	var (
		slice = []int{1, 2, 3, 4, 5}
		array = [4]int{}
	)

	output.Stdoutl("[slice]", slice)
	output.Stdoutl("[array]", array)
	output.StdoutHr()

	// copy(array[:], slice[:len(array)])としても問題はないが
	// copy() は、コピーする要素数を指定した２つの引数の少ない方で
	// 決定してコピーしてくれるので、以下で良い
	copy(array[:], slice)

	output.Stdoutl("[slice]", slice)
	output.Stdoutl("[array]", array)

	return nil
}

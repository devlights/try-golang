package effectivego17

import "github.com/devlights/gomy/output"

// Append -- Effective Go - Append の 内容についてのサンプルです。
func Append() error {
	/*
		https://golang.org/doc/effective_go.html#append

		- append() は 組み込み関数
		- スライスに要素を追加する際に利用する
		- append() は 戻り値で要素が追加された新たなスライスを返す
	*/
	// 基本パターン
	x := []int{1, 2, 3}
	x = append(x, 4, 5)

	output.Stdoutl("(1)", x)

	// 特定のスライスに対して、別のスライスの要素を全部追加
	y := []int{6, 7, 8}
	x = append(x, y...)

	output.Stdoutl("(2)", x)

	// スライスをコピー
	// 組み込みの copy() を利用するのが最もベーシックであるが
	// 空のスライスを用意して、全部 append() しても同じこと
	z := make([]int, 0)
	z = append(z, x...)

	output.Stdoutl("(3)", z)

	return nil
}

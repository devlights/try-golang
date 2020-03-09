package slices

import "github.com/devlights/try-golang/output"

// SliceClear は、スライスのクリア、及び、nilスライスと空のスライスについてのサンプルです.
func SliceClear() error {
	// ----------------------------------------------------------------
	// スライスのクリア、及び、nilスライスと空のスライスについて
	//
	// スライスで以下の２つは機能的に同じ
	// (1) nilスライス
	//     var s []int
	// (2) non-nilだけど長さが0のスライス（空のスライス)
	//     var s []int{}
	//
	// 上記のどちらも len(s), cap(s) は 0 となる。
	// Go Wiki では、(1) の nilスライスを推奨している。
	// ただし、JSONを扱う場合、nilスライスはnullとなってしまい
	// 空スライスは [] と表現されるので、JSONの場合は空スライスの方が良い。
	//
	// 個人的には、他の言語のクセで []int{} の方が好きなので、こっちをよく使ってしまっている..
	// Goらしく、nilスライスを使うようにしようと矯正中。
	// 余談であるが、nilスライスにしておくと、範囲外のアクセスをしているコードがある場合に
	// GoLandだとインスペクションで警告してくれるので、こっちの方が良さそう。
	// (空スライスだとインスペクションで引っかからない)
	//
	// スライスのクリアには、大きく２つある。
	// (1) nilを代入
	//     nilスライスになるので、データがクリアされる。
	//     スライスに設定されていたデータは、GCによって開放される対象となる
	// (2) スライスが参照している場所を変更する
	//     s[:0]と代入することで、長さが0の範囲を参照するスライスにしてしまう。
	//     この場合、参照している範囲を変えただけなのでメモリはクリアされていない
	//     見た目上、長さは0になるが、範囲を広げると元のデータが見える。
	//     capはそのまま。
	//
	// REF:
	//   https://github.com/golang/go/wiki/CodeReviewComments#declaring-empty-slices
	//   https://yourbasic.org/golang/clear-slice/
	// ----------------------------------------------------------------
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{1, 2, 3, 4, 5}
	s3 := []int{1, 2, 3, 4, 5}

	output.Stdoutl("s1", s1, len(s1), cap(s1))
	output.Stdoutl("s2", s2, len(s2), cap(s2))
	output.Stdoutl("s3", s3, len(s3), cap(s3))

	// nilを設定
	// len(s1), cap(s1) も 0 になる
	s1 = nil
	output.Stdoutl("s1", s1, len(s1), cap(s1))

	// 範囲を0にする
	// len(s2) は 0 になるが cap は残る
	s2 = s2[:0]
	output.Stdoutl("s2", s2, len(s2), cap(s2))

	// 空スライスを設定
	// len(), cap() ともに 0 となる
	s3 = []int{}
	output.Stdoutl("s3", s3, len(s3), cap(s3))

	// メモリ上にデータは残っているので範囲を広げると
	// 元のデータは見える
	output.Stdoutl("s2[:2]", s2[:2])

	// nil スライスの場合は、メモリからデータを消しているので
	// s1[:2]とかすると存在しない範囲に対してのアクセスとなる
	// つまり panic する
	defer func() {
		err := recover()
		output.Stdoutl("s1[:2]", err)
	}()

	//noinspection GoNilness
	output.Stdoutl("s1[:2]", s1[:2])
	// nil スライスと機能的には同じなので、以下は同じく panic する
	// nil スライスだとGoLandのインスペクションで警告されるが
	// 空スライスだと警告されない。
	output.Stdoutl("s3[:2]", s3[:2])

	return nil
}

package slices

import (
	"github.com/devlights/gomy/output"
)

// RemoveAllElements -- スライスの全要素を削除するサンプルです.
//
// REFERENCES:
//   - https://yourbasic.org/golang/clear-slice/
func RemoveAllElements() error {
	// -------------------------------------------------------
	// スライスから全要素を削除するには シンプル に nil を代入する.
	//
	// 他に参照が存在しない場合は GC によりメモリから削除される.
	// -------------------------------------------------------
	s1 := []int{1, 2, 3}
	output.Stdoutf("[before]", "%v\tlen=%d\tcap=%d\n", s1, len(s1), cap(s1))

	s1 = nil
	output.Stdoutf("[after]", "%v\tlen=%d\tcap=%d\n", s1, len(s1), cap(s1))

	// 新たなスライスを得ようとしても、以前のスライスは nil を設定しているので、もう存在しない
	// （つまり、cap も 0 となっている)
	//
	// なので、以下は panic する
	/*
		s2 := s1[:1]
		output.Stdoutl("s1[:1]", s2)
	*/
	// panic: runtime error: slice bounds out of range [:1] with capacity 0

	return nil
}

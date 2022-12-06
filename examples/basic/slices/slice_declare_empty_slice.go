package slices

import "github.com/devlights/gomy/output"

// DeclareEmtpySlice -- 空のスライスを宣言する際のお作法についてのサンプルです。
//
// # REFERENCES
//   - https://github.com/golang/go/wiki/CodeReviewComments#declaring-empty-slices
func DeclareEmtpySlice() error {
	var (
		s1 []int     // nil スライス
		s2 = []int{} // non-nil で 要素が 0 のスライス
	)

	// どちらも len=0, cap=0
	output.Stdoutf("[S1]", "len=%v\tcap=%v\n", len(s1), cap(s1))
	output.Stdoutf("[S2]", "len=%v\tcap=%v\n", len(s2), cap(s2))

	// 上の s1, s2 はともに同じ意味となる
	//
	// Go では、 上記のように空スライスを表現する場合 nil スライス の方が好ましいとされている。
	// （JSONでシリアライズする場合などの特別な場合を除いて）
	//
	// https://github.com/golang/go/wiki/CodeReviewComments#declaring-empty-slices を参照。

	return nil
}

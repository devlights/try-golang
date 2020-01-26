package sort_

import (
	"github.com/devlights/try-golang/lib/output"
	"sort"
)

type (
	Sequence []int
)

// sort.Interface の 実装
func (s Sequence) Len() int {
	return len(s)
}

// sort.Interface の 実装
func (s Sequence) Less(i, j int) bool {
	return s[i] < s[j]
}

// sort.Interface の 実装
func (s Sequence) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func SortInterface() error {
	// -------------------------------------------
	// sort.Sort() について
	//
	// sort.Sort() を呼び出すためには
	// 以下のインターフェースを実装している必要がある
	// type Interface interface {
	//     Len() int
	//     Less(i, j int) bool
	//     Swap(i, j int)
	// }
	//
	// Lenは、対象シーケンスの要素数を返す処理を記述する.
	// Lessは、比較結果を記述する。名前の通り「小さい」場合の比較処理を書く.
	// Swapは、指定された要素を入れ替えるための処理を記述する.
	// -------------------------------------------
	var (
		s = Sequence{66, 1, 77, 773, 87, 32}
	)

	// Sequence型は、sort.Interfaceを実装しているので
	// sort.Sort()に渡す事ができる
	output.Stdoutf("(1)", "before: %v\n", s)
	sort.Sort(s)
	output.Stdoutf("(2)", "after : %v\n", s)

	return nil
}

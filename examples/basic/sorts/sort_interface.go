package sorts

import (
	"sort"

	"github.com/devlights/gomy/output"
)

type (
	sequence []int
)

// Len -- sort.Interface の 実装
func (s sequence) Len() int {
	return len(s)
}

// Less -- sort.Interface の 実装
func (s sequence) Less(i, j int) bool {
	return s[i] < s[j]
}

// Swap -- sort.Interface の 実装
func (s sequence) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// SortInterface -- sort.Sort() の際に必要となるインターフェースについてのサンプルです。
func SortInterface() error {
	// -------------------------------------------
	// sort.Sort(), sort.Stable(), sort.Reverse() について
	//
	// 上記３つの関数はどれも同じ引数を要求する
	//   sort.Sort(data Interface)
	//   sort.Stable(data Interface)
	//   sort.Reverse(data Interface)
	// 呼び出すためには
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
	//
	// sort.Sort() は 非安定ソート
	// sort.Stable() は 安定ソート
	// となっている。
	//
	// sort.Reverse() は、少し癖がある関数。
	// この関数は、実際に降順ソートを行うのではなくて、引数で指定されたデータを
	// 降順ソートするためのInterfaceを返してくれる関数。
	// なので、sort.Reverseを呼んだだけでは データはソートされない。
	// 以下のように sort.Reverse で得られたデータを sort.Sort に与えることで
	// 降順ソートされる。
	//
	//   sort.Sort(sort.Reverse(x))
	//
	// -------------------------------------------
	var (
		s1 = sequence{66, 1, 77, 773, 87, 32}
		s2 = sequence{66, 1, 77, 773, 87, 32}
		s3 = sequence{66, 1, 77, 773, 87, 32}
	)

	// Sequence型は、sort.Interfaceを実装しているので
	// sort.Sort()に渡す事ができる
	output.Stdoutf("sort.Sort", "before: %v\n", s1)
	sort.Sort(s1)
	output.Stdoutf("sort.Sort", "after : %v\n", s1)

	output.Stdoutf("sort.Stable", "before: %v\n", s2)
	sort.Stable(s2)
	output.Stdoutf("sort.Stable", "after : %v\n", s2)

	output.Stdoutf("sort.Reverse", "before: %v\n", s3)
	sort.Sort(sort.Reverse(s3))
	output.Stdoutf("sort.Reverse", "after : %v\n", s3)

	return nil
}

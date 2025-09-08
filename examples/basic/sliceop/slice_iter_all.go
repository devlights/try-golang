package sliceop

import (
	"iter"
	"maps"
	"slices"

	"github.com/devlights/gomy/output"
)

// IterAll は、Go 1.23で追加された slices.All() のサンプルです。
//
// > All returns an iterator over index-value pairs in the slice in the usual order.
// > (Allは、スライス要素をインデックス付きで順番に返すイテレータを返します。)
//
// 一番シンプルな形のイテレータ関数。元のスライスをそのままインデックス付きのイテレータにしてくれるだけ。
// iter.Seq2[int, E] を要求する他の関数を呼び出す際などに経由させて使える。
//
// # REFERENCES
//   - https://pkg.go.dev/slices@go1.23.4#All
func IterAll() error {
	var (
		s  []string
		it iter.Seq2[int, string]
	)

	s = []string{"golang", "javascript", "java", "csharp", "python", "rust"}
	it = slices.All(s)

	for i, v := range it {
		output.Stdoutf("[i, v]", "%d, %v\n", i, v)
	}

	output.StdoutHr()

	// maps.Collect() は、iter.Seq2[K,V] を受け取り、map[K]V にして返してくれる関数
	var (
		m = maps.Collect(it)
	)
	output.Stdoutl("[key=3]", m[3])

	return nil
}

package slices

import (
	"iter"
	"slices"

	"github.com/devlights/gomy/output"
)

// IterValues は、Go 1.23で追加された slices.Values() のサンプルです。
//
// > Values returns an iterator that yields the slice elements in order.
// > (Valuesは、スライス要素を順番に返すイテレータを返します。)
//
// 一番シンプルな形のイテレータ関数。元のスライスをそのままイテレータにしてくれるだけ。
// iter.Seq[E] を要求する他の関数を呼び出す際などに経由させて使える。
//
// # REFERENCES
//   - https://pkg.go.dev/slices@go1.23.4#Values
func IterValues() error {
	var (
		s []int
		i iter.Seq[int]
	)

	s = []int{3, 5, 2, 1, 4}
	i = slices.Values(s)

	for v := range i {
		output.Stdoutl("[i]", v)
	}

	output.StdoutHr()

	// slices.Sorted()もGo1.23で追加された関数。iter.Seq[E]を受け取る.
	for v := range slices.Sorted(i) {
		output.Stdoutl("[i]", v)
	}
	return nil
}

package sorts

import (
	"sort"

	"github.com/devlights/gomy/output"
)

// SortSliceStable -- スライスのソートについてのサンプルです. (stable sort)
func SortSliceStable() error {
	// -----------------------------------------------
	// スライスのソートについて
	//
	// スライスをソートする場合
	//   sort.Slice(interface{}, func(i,j int) bool)
	//   or
	//   sort.SliceStable(interface{}, func(i,j int) bool)
	// を利用する.
	//
	// sort.Sliceの方は unstable sort
	// sort.SliceStableの方は stable sort
	// となっている。
	//
	// 注意点として、上記の関数を呼び出すと元のスライスが
	// 書き換わってしまうことに注意。
	//
	// どちらの関数も第一引数が interface{} となっているが
	// スライス以外を渡すと panic する。
	// -----------------------------------------------
	var (
		sli1 = make([]int, 0, 0)
	)

	for i := 10; i > 0; i-- {
		sli1 = append(sli1, i)
	}

	output.Stdoutl("before", sli1)

	// sort.SliceStable は 安定ソート、つまり、同値の場合に元の順序がキープされる
	sort.SliceStable(sli1, func(i, j int) bool {
		return sli1[i] < sli1[j]
	})

	output.Stdoutl("after", sli1)

	return nil
}

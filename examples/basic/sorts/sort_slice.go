package sorts

import (
	"sort"

	"github.com/devlights/gomy/output"
)

// SortSliceUnStable -- スライスのソートについてのサンプルです. (unstable sort)
func SortSliceUnStable() error {
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
		sli1 = make([]int, 0)
	)

	for i := 10; i > 0; i-- {
		sli1 = append(sli1, i)
	}

	output.Stdoutl("before", sli1)

	// sort.Slice は 非安定ソート
	sort.Slice(sli1, func(i, j int) bool {
		return sli1[i] < sli1[j]
	})

	output.Stdoutl("after", sli1)

	// sort.Slice にスライス以外を渡すと panic する
	// 参照: https://blog.golang.org/defer-panic-and-recover
	defer func() {
		if r := recover(); r != nil {
			output.Stdoutl("recover", r)
		}
	}()

	/* 以下を実行すると panic する
	s := "helloworld"
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	*/

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: sort_slice_unstable

	   [Name] "sort_slice_unstable"
	   before               [10 9 8 7 6 5 4 3 2 1]
	   after                [1 2 3 4 5 6 7 8 9 10]


	   [Elapsed] 25.53µs
	*/

}

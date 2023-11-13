package array

import (
	"sort"

	"github.com/devlights/gomy/iter"
	"github.com/devlights/gomy/output"
)

// MultiDemension -- Goで2次元以上の配列（スライス）を利用する方法についてのサンプルです.
//
// # REFERENCES
//   - https://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#one_dim_slice_arr
func MultiDemension() error {
	const (
		x = 2
		y = 10
	)

	var (
		a = make([][]int, x)
	)

	for i := range iter.Range(x) {
		a[i] = make([]int, y)

		for j := range iter.Range(y) {
			a[i][j] = 99 - j
		}
	}

	sort.Slice(a[0], func(i, j int) bool {
		return !(i < j)
	})

	output.Stdoutl("[multi-dementional]", a)

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: array_multi_demention

	   [Name] "array_multi_demention"
	   [multi-dementional]  [[90 91 92 93 94 95 96 97 98 99] [99 98 97 96 95 94 93 92 91 90]]


	   [Elapsed] 68.52µs
	*/

}

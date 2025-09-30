package main

import (
	"slices"
	"testing"
)

const (
	NumItems = 5 * 1024 * 1024
)

func voidfn(_ []int) {}

// 方法1: 事前に容量を確保してからappendする方法
func BenchmarkPreAllocatedAppend(b *testing.B) {
	var (
		items = make([]int, 0, NumItems)
	)
	for i := range NumItems {
		items = append(items, i)
	}

	for b.Loop() {
		var (
			dst = make([]int, 0, NumItems)
		)
		dst = append(dst, items...)
		voidfn(dst)
	}
}

// 方法2: 空スライスに直接appendする方法
func BenchmarkEmptySliceAppend(b *testing.B) {
	var (
		items = make([]int, 0, NumItems)
	)
	for i := range NumItems {
		items = append(items, i)
	}

	for b.Loop() {
		// 空スライスに直接append
		dst := append([]int{}, items...)
		voidfn(dst)
	}
}

// 方法3: 容量をリセットしてからappendする方法
func BenchmarkCapacityResetAppend(b *testing.B) {
	var (
		items = make([]int, 0, NumItems)
	)
	for i := range NumItems {
		items = append(items, i)
	}

	for b.Loop() {
		var (
			dst = make([]int, 0, NumItems)
		)
		// スライスの長さと容量を0にリセットしてからappend
		dst = append(dst[:0:0], items...)
		voidfn(dst)
	}
}

// 方法4: 従来のcopy()を使った方法
func BenchmarkBuiltinCopy(b *testing.B) {
	var (
		items = make([]int, 0, NumItems)
	)
	for i := range NumItems {
		items = append(items, i)
	}

	for b.Loop() {
		dst := make([]int, len(items))
		copy(dst, items)
		voidfn(dst)
	}
}

// 方法5: slices.Clone()を使った方法
func BenchmarkSlicesClone(b *testing.B) {
	var (
		items = make([]int, 0, NumItems)
	)
	for i := range NumItems {
		items = append(items, i)
	}

	for b.Loop() {
		dst := slices.Clone(items)
		voidfn(dst)
	}
}

/*
$ task
task: [default] go test -bench . -benchmem -count 3 | tee bench.txt
goos: linux
goarch: amd64
pkg: github.com/devlights/try-golang/examples/benchmarks/03.slice_clone
cpu: Intel(R) Core(TM) Ultra 5 125H
BenchmarkPreAllocatedAppend-18               134           8116946 ns/op        41943090 B/op          1 allocs/op
BenchmarkPreAllocatedAppend-18               150           7897353 ns/op        41943049 B/op          1 allocs/op
BenchmarkPreAllocatedAppend-18               153           7697424 ns/op        41943048 B/op          1 allocs/op
BenchmarkEmptySliceAppend-18                 234           5234496 ns/op        41943072 B/op          1 allocs/op
BenchmarkEmptySliceAppend-18                 246           4907417 ns/op        41943081 B/op          1 allocs/op
BenchmarkEmptySliceAppend-18                 174           6013573 ns/op        41943054 B/op          1 allocs/op
BenchmarkCapacityResetAppend-18              153           7743873 ns/op        83886086 B/op          2 allocs/op
BenchmarkCapacityResetAppend-18              152           7804002 ns/op        83886092 B/op          2 allocs/op
BenchmarkCapacityResetAppend-18              150           7891077 ns/op        83886088 B/op          2 allocs/op
BenchmarkBuiltinCopy-18                      151           7796522 ns/op        41943046 B/op          1 allocs/op
BenchmarkBuiltinCopy-18                      152           7813188 ns/op        41943050 B/op          1 allocs/op
BenchmarkBuiltinCopy-18                      158           7525531 ns/op        41943050 B/op          1 allocs/op
BenchmarkSlicesClone-18                      255           5556701 ns/op        41943047 B/op          1 allocs/op
BenchmarkSlicesClone-18                      187           6222304 ns/op        41943047 B/op          1 allocs/op
BenchmarkSlicesClone-18                      204           5850473 ns/op        41943044 B/op          1 allocs/op
PASS
ok      github.com/devlights/try-golang/examples/benchmarks/03.slice_clone      17.955s
task: [default] benchstat bench.txt
goos: linux
goarch: amd64
pkg: github.com/devlights/try-golang/examples/benchmarks/03.slice_clone
cpu: Intel(R) Core(TM) Ultra 5 125H
                       │  bench.txt   │
                       │    sec/op    │
PreAllocatedAppend-18    7.897m ± ∞ ¹
EmptySliceAppend-18      5.234m ± ∞ ¹
CapacityResetAppend-18   7.804m ± ∞ ¹
BuiltinCopy-18           7.797m ± ∞ ¹
SlicesClone-18           5.850m ± ∞ ¹
geomean                  6.816m
¹ need >= 6 samples for confidence interval at level 0.95

                       │   bench.txt   │
                       │     B/op      │
PreAllocatedAppend-18    40.00Mi ± ∞ ¹
EmptySliceAppend-18      40.00Mi ± ∞ ¹
CapacityResetAppend-18   80.00Mi ± ∞ ¹
BuiltinCopy-18           40.00Mi ± ∞ ¹
SlicesClone-18           40.00Mi ± ∞ ¹
geomean                  45.95Mi
¹ need >= 6 samples for confidence interval at level 0.95

                       │  bench.txt  │
                       │  allocs/op  │
PreAllocatedAppend-18    1.000 ± ∞ ¹
EmptySliceAppend-18      1.000 ± ∞ ¹
CapacityResetAppend-18   2.000 ± ∞ ¹
BuiltinCopy-18           1.000 ± ∞ ¹
SlicesClone-18           1.000 ± ∞ ¹
geomean                  1.149
¹ need >= 6 samples for confidence interval at level 0.95
task: [default] rm -f bench.txt
*/

/*
	結果として、何度試しても空スライスにAppendする方法が最も速かった。
	つまり、
		dst := append([]int{}, items...)
	が速い。
	slices.Clone()は、内部で上記を行っている
		// Avoid s[:0:0] as it leads to unwanted liveness when cloning a
		// zero-length slice of a large array; see https://go.dev/issue/68488.
		return append(S{}, s...)
*/

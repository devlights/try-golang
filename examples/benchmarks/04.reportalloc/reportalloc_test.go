package main

import (
	"slices"
	"testing"
)

func voidfn(_ []int) {}

func BenchmarkWithReportAllocs(b *testing.B) {
	src := make([]int, 255)
	for i := range src {
		src[i] = i
	}

	// -benchmem オプションを付与したのと同じことになる
	// オプションとして指定すると全ベンチマークがONとなるが
	// ReportAllocs()の場合は、呼び出したベンチマークのみがONとなる。
	b.ReportAllocs()

	for b.Loop() {
		dst := slices.Clone(src)
		voidfn(dst)
	}
}

func BenchmarkWithoutReportAllocs(b *testing.B) {
	src := make([]int, 255)
	for i := range src {
		src[i] = i
	}

	for b.Loop() {
		dst := append([]int{}, src...)
		voidfn(dst)
	}
}

/*
$ task
task: [default] go test . -bench .
goos: linux
goarch: amd64
pkg: github.com/devlights/try-golang/examples/benchmarks/04.reportalloc
cpu: Intel(R) Core(TM) Ultra 5 125H
BenchmarkWithReportAllocs-18             3909057               310.4 ns/op          2048 B/op          1 allocs/op
BenchmarkWithoutReportAllocs-18          3827202               322.6 ns/op
PASS
ok      github.com/devlights/try-golang/examples/benchmarks/04.reportalloc      2.452s

task: [default] go test . -bench . -benchmem
goos: linux
goarch: amd64
pkg: github.com/devlights/try-golang/examples/benchmarks/04.reportalloc
cpu: Intel(R) Core(TM) Ultra 5 125H
BenchmarkWithReportAllocs-18             3708218               322.0 ns/op          2048 B/op          1 allocs/op
BenchmarkWithoutReportAllocs-18          3210252               338.4 ns/op          2048 B/op          1 allocs/op
PASS
ok      github.com/devlights/try-golang/examples/benchmarks/04.reportalloc      2.285s
*/

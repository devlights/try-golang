package main

import (
	"fmt"
	"testing"
)

// Examplefib -- fib() の Exampleテスト
func Examplefib() {
	fmt.Println(fib(40))

	// Output:
	// 102334155
}

// TestFib -- fib() の ユニットテスト
func TestFib(t *testing.T) {
	tests := []struct {
		in, want int
	}{
		{1, 1},
		{2, 1},
		{3, 2},
		{10, 55},
		{40, 102334155},
		{41, 165580141},
		{42, 267914296},
		{43, 433494437},
	}

	for _, test := range tests {
		test := test
		title := fmt.Sprintf("fib:%d", test.in)

		t.Run(title, func(t *testing.T) {
			t.Parallel()

			got := fib(test.in)
			if test.want != got {
				t.Errorf("[want] %v\t[got] %v", test.want, got)
			}
		})
	}
}

// BenchmarkAppend_WithoutLen -- ベンチマークのサンプル
func BenchmarkAppend_WithoutLen(b *testing.B) {
	s := make([]string, 0)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		//lint:ignore SA4010 It's ok because this is just a example.
		s = append(s, fmt.Sprintf("%d", i))
	}
}

// BenchmarkAppend_WithLen -- ベンチマークのサンプル
func BenchmarkAppend_WithLen(b *testing.B) {
	s := make([]string, 0, b.N)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		//lint:ignore SA4010 It's ok because this is just a example.
		s = append(s, fmt.Sprintf("%d", i))
	}
}

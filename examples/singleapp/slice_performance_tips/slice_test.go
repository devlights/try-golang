package main

import (
	"strconv"
	"testing"
)

func BenchmarkSliceLen0Append(b *testing.B) {
	var (
		s []string
	)

	for i := 0; i < b.N; i++ {
		//lint:ignore SA4010 ok
		s = append(s, strconv.Itoa(i))
	}
}

func BenchmarkSliceLenN(b *testing.B) {
	var (
		s = make([]string, b.N)
	)

	for i := 0; i < b.N; i++ {
		s[i] = strconv.Itoa(i)
	}
}

func BenchmarkSliceLenNAppend(b *testing.B) {
	var (
		s = make([]string, 0, b.N)
	)

	for i := 0; i < b.N; i++ {
		//lint:ignore SA4010 ok
		s = append(s, strconv.Itoa(i))
	}
}

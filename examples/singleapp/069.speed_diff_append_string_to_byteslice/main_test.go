package main

import (
	"fmt"
	"strconv"
	"testing"
)

func BenchmarkUseFmtSprintf(b *testing.B) {
	b.StopTimer()
	buf := make([]byte, 0, 1024*1024*500)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		s := fmt.Sprintf("%s:%d\n", "key", i)
		buf = append(buf, s...) //lint:ignore SA4010 It's ok.
	}
}

func BenchmarkUseFmtAppendf(b *testing.B) {
	b.StopTimer()
	buf := make([]byte, 0, 1024*1024*500)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		buf = fmt.Appendf(buf, "%s:%d\n", "key", i)
	}
}

func BenchmarkUseDirectAppend(b *testing.B) {
	b.StopTimer()
	buf := make([]byte, 0, 1024*1024*500)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		buf = append(buf, "key:"...)          //lint:ignore SA4010 It's ok.
		buf = append(buf, strconv.Itoa(i)...) //lint:ignore SA4010 It's ok.
		buf = append(buf, '\n')               //lint:ignore SA4010 It's ok.
	}
}

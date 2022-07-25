package main

import (
	"bytes"
	"strings"
	"testing"
)

func BenchmarkBufferConcatenate(b *testing.B) {
	var (
		buf bytes.Buffer
	)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		buf.WriteString("i")
	}
	b.StopTimer()

	b.Log(buf.Len())
}

func BenchmarkStringBuilderConcatenate(b *testing.B) {
	var (
		sb strings.Builder
	)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sb.WriteString("i")
	}
	b.StopTimer()

	b.Log(sb.Len())
}

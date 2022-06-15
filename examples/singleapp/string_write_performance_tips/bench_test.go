// REFERENCES
//   - https://stackoverflow.com/questions/1760757/how-to-efficiently-concatenate-strings-in-go/47798475#47798475
//   -

package main

import (
	"bytes"
	"strings"
	"testing"
)

func BenchmarkStringBuilder(b *testing.B) {
	var (
		buf strings.Builder
	)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf.WriteString("x")
	}
	b.StopTimer()

	if buf.Len() != b.N {
		b.Errorf("[want] %v\t[got] %v", b.N, buf.Len())
	}
}

func BenchmarkByteBuffer(b *testing.B) {
	var (
		buf bytes.Buffer
	)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf.WriteString("x")
	}
	b.StopTimer()

	if buf.Len() != b.N {
		b.Errorf("[want] %v\t[got] %v", b.N, buf.Len())
	}
}

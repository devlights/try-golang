package main

import (
	"testing"

	"github.com/nalgeon/be"
)

func sum(x ...int) int64 {
	var (
		total int64
	)
	for _, v := range x {
		total += int64(v)
	}

	return total
}

func TestBeEqual(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		be.Equal(t, sum(1, 2, 3), int64(6))
	})
	t.Run("ng", func(t *testing.T) {
		be.Equal(t, sum(1, 2, 3), int64(10))
	})
	t.Run("ok2", func(t *testing.T) {
		// wantsには複数指定できる。複数していした場合はどれかが合致すればOKとなる。
		be.Equal(t, sum(1, 2, 3), int64(10), int64(6), int64(99))
	})
}

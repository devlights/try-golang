package main

import (
	"errors"
	"fmt"
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

func errSometime(v int) error {
	switch {
	case v > 100:
		return fmt.Errorf("invalid value: %d", v)
	case v > 10:
		return errors.New("invalid value")
	default:
		return nil
	}
}

func TestBeTrue(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		be.True(t, sum(1, 2, 3) == 6)
	})
	t.Run("ng", func(t *testing.T) {
		be.True(t, sum(1, 2, 3) != 6)
	})
	t.Run("ok2", func(t *testing.T) {
		be.True(t, errSometime(1) == nil)
	})
	t.Run("ng2", func(t *testing.T) {
		be.True(t, errSometime(111) == nil)
	})
}

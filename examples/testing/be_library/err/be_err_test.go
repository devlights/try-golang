package main

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/nalgeon/be"
)

type (
	ValueError struct {
		V int
	}
)

var _ error = (*ValueError)(nil)

func (v *ValueError) Error() string {
	return fmt.Sprintf("invalid value: %d", v.V)
}

var (
	ErrInvalid = errors.New("invalid value")
)

func errSometime(v int) error {
	switch {
	case v > 100:
		return &ValueError{v}
	case v > 10:
		return ErrInvalid
	default:
		return nil
	}
}

func TestBeErr(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		be.Err(t, errSometime(1), nil)
	})
	t.Run("ng", func(t *testing.T) {
		be.Err(t, errSometime(99), nil)
	})
	t.Run("return-nil", func(t *testing.T) {
		err := errSometime(5)
		be.Err(t, err, nil)
	})
	t.Run("return-err", func(t *testing.T) {
		err := errSometime(99)
		be.Err(t, err)
	})
	t.Run("check-message", func(t *testing.T) {
		err := errSometime(20)
		be.Err(t, err, "invalid") // 指定したメッセージが含まれていればOKとなる
	})
	t.Run("check-type-is", func(t *testing.T) {
		err := errSometime(20)
		be.Err(t, err, ErrInvalid)
	})
	t.Run("check-type-as", func(t *testing.T) {
		err := errSometime(199)
		be.Err(t, err, reflect.TypeFor[*ValueError]()) // Asの場合は reflect.Type を指定する
	})
}

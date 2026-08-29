// ジェネリックメソッドのサンプル
//
// # REFERENCES
//   - https://devlights.hatenablog.com/entry/2026/07/03/073000
//   - https://devlights.hatenablog.com/entry/2026/08/18/073000
package main

import (
	"fmt"
	"strconv"
)

type (
	Item[T any] struct {
		value T
	}
)

func NewItem[T any](v T) *Item[T] {
	return &Item[T]{v}
}

// Convert は、ジェネリックメソッドのサンプルです。
func (i *Item[T]) Convert[U any](fn func(T) U) *Item[U] {
	return &Item[U]{value: fn(i.value)}
}

func main() {
	i1 := NewItem(100)
	i2 := i1.Convert[string](func(i int) string {
		return strconv.Itoa(i)
	})

	fmt.Printf("i1: %T\ti2:%T\n", i1, i2)
}

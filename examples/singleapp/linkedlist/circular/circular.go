package main

import (
	"fmt"
	"strings"
)

type (
	// Node は、連結リストのノートを表します.
	Node[T any] struct {
		Value T        // 値
		Next  *Node[T] // 次要素
	}

	// Circular は、循環連結リストを表します.
	// Capacityを超えた場合、最も古い要素から上書きされます.
	Circular[T any] struct {
		Head     *Node[T] // 先頭
		Tail     *Node[T] // 末尾
		Size     int      // 要素数
		Capacity int      // キャパシティ
	}
)

// String は、Nodeの文字列表現を返します.
func (me *Node[T]) String() string {
	if me == nil {
		return ""
	}

	if me.Next == nil {
		return fmt.Sprintf("(v:%v,n:nil)", me.Value)
	}

	return fmt.Sprintf("(v:%v,n:%v)", me.Value, me.Next.Value)
}

// NewCircular は、指定されたキャパシティを持つ[*Circular]を生成します.
func NewCircular[T any](capacity int) *Circular[T] {
	if capacity <= 0 {
		panic("capacity must be positive")
	}

	return &Circular[T]{
		Head:     nil,
		Tail:     nil,
		Size:     0,
		Capacity: capacity,
	}
}

// Add は、末尾に新しい要素を追加します.
func (me *Circular[T]) Add(value T) {
	var (
		n = &Node[T]{Value: value}
	)

	if me.Head == nil {
		me.Head = n
		me.Tail = n
		me.Size = 1

		return
	}

	if me.Size < me.Capacity {
		me.Tail.Next = n
		me.Tail = n
		me.Size++
	} else {
		me.Head = me.Head.Next
		me.Tail.Next = n
		me.Tail = n
	}
}

// Delete は、指定した値に合致する最初のノードを削除します.
//
// 型パラメータでは、comparableかどうかが判別できませんので、比較判定関数を指定する必要があります。
func (me *Circular[T]) Delete(value T, equal func(v1, v2 T) bool) bool {
	if me.Head == nil {
		return false
	}

	if equal(me.Head.Value, value) {
		me.Head = me.Head.Next
		me.Size--

		if me.Size == 0 {
			me.Tail = nil
		}

		return true
	}

	var (
		current = me.Head
	)
	for current.Next != nil {
		if equal(current.Next.Value, value) {
			current.Next = current.Next.Next // 削除対象を抜いて要素を詰める

			if current.Next == nil {
				me.Tail = current
			}

			me.Size--

			return true
		}

		current = current.Next
	}

	return false
}

// ToSlice は、要素の値をスライスにして返します.
func (me *Circular[T]) ToSlice() []T {
	var (
		result  = make([]T, 0, me.Size)
		current = me.Head
	)

	for range me.Size {
		result = append(result, current.Value)
		current = current.Next
	}

	return result
}

// String は、[*Circular]の文字列表現を返します.
func (me *Circular[T]) String() string {
	if me.Head == nil {
		return "[]"
	}

	var (
		sb      strings.Builder
		current = me.Head
	)

	sb.WriteString("[")
	{
		for i := 0; i < me.Size; i++ {
			if i > 0 {
				sb.WriteString(" -> ")
			}

			sb.WriteString(fmt.Sprintf("%v", current.Value))
			current = current.Next
		}
	}
	sb.WriteString(" ->]")

	return sb.String()
}

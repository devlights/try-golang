// Package ring は、container/ring/Ring をジェネリックにしたものが配置されています。
package ring

import "container/ring"

// Ring[T] は、container/ring/Ring のジェネリック版です.
//
// 標準の *ring.Ring と異なり、値の設定は SetValue() で行います。
type Ring[T any] struct {
	r *ring.Ring // 実際の *ring.Ring
	v T          // 値
}

// New は、新しい *Ring[T] を指定した容量で作成します。
func New[T any](n int) *Ring[T] {
	var v T
	return &Ring[T]{ring.New(n), v}
}

// SetValue は、値を設定します。
func (me *Ring[T]) SetValue(v T) {
	me.v = v
	me.r.Value = v
}

// Do は、標準ライブラリの container/ring/Ring.Do と同じ動きをします。
//
// # REFERENCES
//   - https://pkg.go.dev/container/ring@go1.21.4#Ring.Do
func (me *Ring[T]) Do(f func(T)) {
	fn := func(v any) {
		f(v.(T))
	}
	me.r.Do(fn)
}

// Len は、標準ライブラリの container/ring/Ring.Len と同じ動きをします。
//
// # REFERENCES
//   - https://pkg.go.dev/container/ring@go1.21.4#Ring.Len
func (me *Ring[T]) Len() int {
	return me.r.Len()
}

// Link は、標準ライブラリの container/ring/Ring.Link と同じ動きをします。
//
// # REFERENCES
//   - https://pkg.go.dev/container/ring@go1.21.4#Ring.Link
func (me *Ring[T]) Link(s *Ring[T]) *Ring[T] {
	return &Ring[T]{me.r.Link(s.r), me.v}
}

// Move は、標準ライブラリの container/ring/Ring.Move と同じ動きをします。
//
// # REFERENCES
//   - https://pkg.go.dev/container/ring@go1.21.4#Ring.Move
func (me *Ring[T]) Move(n int) *Ring[T] {
	return &Ring[T]{me.r.Move(n), me.v}
}

// Next は、標準ライブラリの container/ring/Ring.Next と同じ動きをします。
//
// # REFERENCES
//   - https://pkg.go.dev/container/ring@go1.21.4#Ring.Next
func (me *Ring[T]) Next() *Ring[T] {
	return &Ring[T]{me.r.Next(), me.v}
}

// Prev は、標準ライブラリの container/ring/Ring.Prev と同じ動きをします。
//
// # REFERENCES
//   - https://pkg.go.dev/container/ring@go1.21.4#Ring.Prev
func (me *Ring[T]) Prev() *Ring[T] {
	return &Ring[T]{me.r.Prev(), me.v}
}

// Unlink は、標準ライブラリの container/ring/Ring.Unlink と同じ動きをします。
//
// # REFERENCES
//   - https://pkg.go.dev/container/ring@go1.21.4#Ring.Unlink
func (me *Ring[T]) Unlink(n int) *Ring[T] {
	return &Ring[T]{me.r.Unlink(n), me.v}
}

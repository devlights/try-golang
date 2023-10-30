// Package map は、カスタムマップを定義しています。
package maps

import "sync"

// SafeMap は、sync.Mapをラップしジェネリックにしたものです。
//
// 内部で sync.Map を内包しています。
//
// # REFERENCES
//   - https://pkg.go.dev/sync@go1.21.3#Map
type SafeMap[K any, V any] struct {
	m sync.Map
}

// Load は、指定したキーに紐づく値を取得します。
func (me *SafeMap[K, V]) Load(key K) (V, bool) {
	v, ok := me.m.Load(key)
	return v.(V), ok
}

// Store は、指定したキーと値を保存します。
func (me *SafeMap[K, V]) Store(key K, value V) {
	me.m.Store(key, value)
}

// Range は、自身をイテレーションし要素毎に fn を呼び出します。
func (me *SafeMap[K, V]) Range(fn func(key K, value V) bool) {
	me.m.Range(func(key any, value any) bool {
		return fn(key.(K), value.(V))
	})
}

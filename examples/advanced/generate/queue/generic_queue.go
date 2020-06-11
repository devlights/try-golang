package queue

//go:generate genny -in=generic_queue.go -out=builtins_queue.go gen "T=bool,string"

import "github.com/cheekybits/genny/generic"

type (
	// T -- キューの型
	T generic.Type
	// TQueue -- キュー
	TQueue struct {
		items []T
	}
)

// NewTQueue -- 新しいキューを生成して返します.
//
//noinspection GoUnusedExportedFunction
func NewTQueue() *TQueue {
	q := new(TQueue)
	q.items = make([]T, 0, 0)
	return q
}

// Count -- データの件数を返します.
func (q *TQueue) Count() int {
	return len(q.items)
}

// Enqueue -- データを投入します.
func (q *TQueue) Enqueue(v T) (ok bool) {
	q.items = append(q.items, v)
	ok = true
	return
}

// Dequeue -- データを取り出します.
func (q *TQueue) Dequeue() (v T, ok bool) {
	if q.Count() == 0 {
		ok = false
		return
	}

	v = q.items[0]
	if q.Count() == 0 {
		ok = true
		return
	}

	q.items = q.items[1:]
	ok = true

	return
}

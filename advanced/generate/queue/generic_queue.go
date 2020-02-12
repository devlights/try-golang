package queue

//go:generate genny -in=generic_queue.go -out=builtins_queue.go gen "T=BUILTINS"

import "github.com/cheekybits/genny/generic"

type (
	T      generic.Type
	TQueue struct {
		items []T
	}
)

//noinspection GoUnusedExportedFunction
func NewTQueue() *TQueue {
	q := new(TQueue)
	q.items = make([]T, 0, 0)
	return q
}

func (q *TQueue) Count() int {
	return len(q.items)
}

func (q *TQueue) Enqueue(v T) (ok bool) {
	q.items = append(q.items, v)
	ok = true
	return
}

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

package stack

//go:generate genny -in=generic_stack.go -out=builtins_stack.go gen "T=bool,string,int"

import "github.com/cheekybits/genny/generic"

type (
	// T -- スタックの型
	T generic.Type
	// TStack -- スタック
	TStack struct {
		items []T
	}
)

// NewTStack -- 新しいスタックを生成して返します.
//
//noinspection GoUnusedExportedFunction
func NewTStack() *TStack {
	v := new(TStack)
	v.items = make([]T, 0, 0)
	return v
}

// Push -- データを投入します.
func (s *TStack) Push(v T) {
	s.items = append(s.items, v)
}

// Pop -- データを取り出します.
func (s *TStack) Pop() T {
	last := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return last
}

package generate

//go:generate genny -in=generic_stack.go -out=builtins_stack.go gen "T=BUILTINS"

import "github.com/cheekybits/genny/generic"

type (
	T      generic.Type
	TStack struct {
		items []T
	}
)

//noinspection GoUnusedExportedFunction
func NewTStack() *TStack {
	v := new(TStack)
	v.items = make([]T, 0, 0)
	return v
}

func (s *TStack) Push(v T) {
	s.items = append(s.items, v)
}

func (s *TStack) Pop() T {
	last := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return last
}

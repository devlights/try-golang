package books

import (
	"github.com/devlights/try-golang/interfaces"
)

type (
	booksExampleRegister struct{}
)

// NewRegister は、books パッケージ用の lib.Register を返します.
func NewRegister() interfaces.Register {
	r := new(booksExampleRegister)
	return r
}

//noinspection GoUnusedParameter
func (r *booksExampleRegister) Regist(m interfaces.ExampleMapping) {
}

package books

import (
	"github.com/devlights/try-golang/mappings"
)

type (
	booksExampleRegister struct{}
)

// NewRegister は、books パッケージ用の lib.Register を返します.
func NewRegister() mappings.Register {
	r := new(booksExampleRegister)
	return r
}

// Regist は、books パッケージ配下に存在するサンプルを登録します.
//noinspection GoUnusedParameter
func (r *booksExampleRegister) Regist(m mappings.ExampleMapping) {
}

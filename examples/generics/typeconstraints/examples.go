package typeconstraints

import (
	"github.com/devlights/try-golang/mapping"
)

type (
	register struct{}
)

// NewRegister は、generics パッケージ用の lib.Register を返します.
func NewRegister() mapping.Register {
	r := new(register)
	return r
}

// Regist は、generics パッケージ配下に存在するサンプルを登録します.
func (r *register) Regist(m mapping.ExampleMapping) {
	m["generics_typeconstraints_typeset"] = TypeSet
	m["generics_typeconstraints_underlyingtype"] = UnderlyingType
}

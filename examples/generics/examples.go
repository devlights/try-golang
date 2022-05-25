package generics

import (
	"github.com/devlights/try-golang/examples/generics/exp_constraints"
	"github.com/devlights/try-golang/examples/generics/exp_maps"
	"github.com/devlights/try-golang/examples/generics/exp_slices"
	"github.com/devlights/try-golang/examples/generics/typeconstraints"
	"github.com/devlights/try-golang/examples/generics/typeparameters"
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
	typeparameters.NewRegister().Regist(m)
	typeconstraints.NewRegister().Regist(m)
	exp_constraints.NewRegister().Regist(m)
	exp_slices.NewRegister().Regist(m)
	exp_maps.NewRegister().Regist(m)
}

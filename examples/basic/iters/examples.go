package iters

import (
	"github.com/devlights/try-golang/mapping"
)

type (
	register struct{}
)

// NewRegister -- このパッケージ用のサンプルを登録する mapping.Register を生成します。
func NewRegister() mapping.Register {
	return new(register)
}

// Regist -- 登録します.
func (r *register) Regist(m mapping.ExampleMapping) {
	m["iters_range_over_func_1"] = Go123RangeOverFunc1
	m["iters_range_over_func_2"] = Go123RangeOverFunc2
	m["iters_range_over_func_3"] = Go123RangeOverFunc3
}

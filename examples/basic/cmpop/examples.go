package cmpop

import "github.com/devlights/try-golang/mapping"

type (
	register struct{}
)

// NewRegister -- このパッケージ用のサンプルを登録する mapping.Register を生成します。
func NewRegister() mapping.Register {
	return new(register)
}

// Regist -- 登録します.
func (r *register) Regist(m mapping.ExampleMapping) {
	m["cmpop_or"] = Or
	m["cmpop_compare"] = Compare
	m["cmpop_composite_sort_keys"] = CompositeSortKeys
}

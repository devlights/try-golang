package types

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
	m["types_basic"] = Basic
	m["types_define_types_easy"] = DefineTypesEasy
	m["types_diff_typealias_and_definedtype"] = DiffTypeAliasAndDefinedType
}

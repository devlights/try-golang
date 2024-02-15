package functions

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
	m["function_one_return_value"] = OneReturnValue
	m["function_multi_return_value"] = MultiReturnValue
	m["function_named_return_value"] = NamedReturnValue
	m["function_blank_identifier"] = BlankIdentifier
}

package functions

import (
	"github.com/devlights/try-golang/mappings"
)

type (
	register struct{}
)

// NewRegister -- このパッケージ用のサンプルを登録する mappings.Register を生成します。
func NewRegister() mappings.Register {
	return new(register)
}

// Regist -- 登録します.
func (r *register) Regist(m mappings.ExampleMapping) {
	m["function_one_return_value"] = FunctionOneReturnValue
	m["function_multi_return_value"] = FunctionMultiReturnValue
	m["function_named_return_value"] = FunctionNamedReturnValue
}

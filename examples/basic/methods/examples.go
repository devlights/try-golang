package methods

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
	m["methods_pointer_or_not"] = PointerOrNot
	m["methods_method_value"] = MethodValue
	m["methods_method_expression"] = MethodExpression
}

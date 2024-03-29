package builtins

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
	m["builtin_print"] = PrintFunc
	m["builtin_minmax"] = MinMax
	m["builtin_clear"] = Clear
	m["builtin_delete"] = Delete
	m["builtin_copy"] = Copy
}

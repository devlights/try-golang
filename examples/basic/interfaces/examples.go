package interfaces

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
	m["interface_basic"] = Basic
	m["interface_composition"] = Composition
	m["interface_ducktyping"] = DuckTyping
	m["interface_verify_compliance"] = VerifyInterfaceCompliance
	m["interface_nil_notnil"] = NilOrNotNil
}

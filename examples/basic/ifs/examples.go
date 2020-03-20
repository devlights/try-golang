package ifs

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

func (r *register) Regist(m mappings.ExampleMapping) {
	m["interface_basic"] = Basic
	m["interface_composition"] = Composition
	m["interface_ducktyping"] = DuckTyping
	m["interface_verify_compliance"] = VerifyInterfaceCompliance
}

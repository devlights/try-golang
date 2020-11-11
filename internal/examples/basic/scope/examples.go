package scope

import (
	"github.com/devlights/try-golang/pkg/mappings"
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
	m["scope_basic"] = Basic
	m["scope_common_mistake1"] = CommonMistake1
	m["scope_common_mistake2"] = CommonMistake2
	m["scope_common_mistake3"] = CommonMistake3
}

package scope

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
	m["scope_basic"] = Basic
	m["scope_common_mistake1"] = CommonMistake1
	m["scope_common_mistake2"] = CommonMistake2
	m["scope_common_mistake3"] = CommonMistake3
}

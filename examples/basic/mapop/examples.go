package mapop

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
	m["mapop_basic"] = MapBasic
	m["mapop_for"] = MapFor
	m["mapop_initialize"] = MapInitialize
	m["mapop_delete"] = MapDelete
	m["mapop_access"] = MapAccess
	m["mapop_deep_equal"] = MapDeepEqual
}

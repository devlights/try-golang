package maps

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
	m["map_basic"] = MapBasic
	m["map_for"] = MapFor
	m["map_initialize"] = MapInitialize
	m["map_delete"] = MapDelete
	m["map_access"] = MapAccess
	m["map_deep_equal"] = MapDeepEqual
}

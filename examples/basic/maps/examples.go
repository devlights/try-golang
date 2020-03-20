package maps

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
	m["map_basic"] = MapBasic
	m["map_for"] = MapFor
	m["map_initialize"] = MapInitialize
	m["map_delete"] = MapDelete
	m["map_access"] = MapAccess
	m["map_deep_equal"] = MapDeepEqual
}

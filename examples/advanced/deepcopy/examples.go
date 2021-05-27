package deepcopy

import "github.com/devlights/try-golang/mapping"

type (
	register struct{}
)

// NewRegister -- このパッケージ用のサンプルを登録する mapping.Register を生成します。
func NewRegister() mapping.Register {
	return &register{}
}

// Regist -- サンプルを登録します。
func (r *register) Regist(m mapping.ExampleMapping) {
	m["deepcopy_gob"] = GobDeepCopy
	m["deepcopy_json"] = JsonDeepCopy
}

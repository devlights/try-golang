package sets

import (
	"github.com/devlights/try-golang/pkg/mappings"
)

type (
	register struct{}
)

// NewRegister -- このパッケージ用のサンプルを登録する mappings.Register を生成します。
func NewRegister() mappings.Register {
	r := new(register)
	return r
}

// Regist -- 登録します.
func (r *register) Regist(m mappings.ExampleMapping) {
	m["set01"] = Set01
	m["set02"] = Set02
	m["set03"] = Set03
	m["set04"] = Set04
	m["set05"] = Set05
	m["mapset_01"] = MapSet01
}

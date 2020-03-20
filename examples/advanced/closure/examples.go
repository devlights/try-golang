package closure

import (
	"github.com/devlights/try-golang/mappings"
)

type (
	register struct{}
)

// NewRegister -- このパッケージ用のサンプルを登録する mappings.Register を生成します。
func NewRegister() mappings.Register {
	r := new(register)
	return r
}

// Regist -- サンプルを登録します。
func (r *register) Regist(m mappings.ExampleMapping) {
	m["closure01"] = Closure01
}

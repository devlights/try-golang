package reflection

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
	m["reflect_typeof"] = TypeOf
}

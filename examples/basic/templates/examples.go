package templates

import (
	"github.com/devlights/try-golang/examples/basic/templates/htmltmpl"
	"github.com/devlights/try-golang/examples/basic/templates/txttmpl"
	"github.com/devlights/try-golang/mapping"
)

type (
	register struct{}
)

// NewRegister -- このパッケージ用のサンプルを登録する mapping.Register を生成します。
func NewRegister() mapping.Register {
	return new(register)
}

// Regist -- 登録します.
func (r *register) Regist(m mapping.ExampleMapping) {
	txttmpl.NewRegister().Regist(m)
	htmltmpl.NewRegister().Regist(m)
}

package advanced

import (
	"github.com/devlights/try-golang/advanced/async"
	"github.com/devlights/try-golang/advanced/closure"
	"github.com/devlights/try-golang/advanced/crypto"
	"github.com/devlights/try-golang/advanced/generate"
	"github.com/devlights/try-golang/advanced/reflection"
	"github.com/devlights/try-golang/advanced/sets"
	"github.com/devlights/try-golang/advanced/xdgspec"
	"github.com/devlights/try-golang/interfaces"
)

type (
	register struct{}
)

// NewRegister は、advanced パッケージ用の lib.Register を返します.
func NewRegister() interfaces.Register {
	r := new(register)
	return r
}

// Regist は、advanced パッケージ配下に存在するサンプルを登録します.
func (r *register) Regist(m interfaces.ExampleMapping) {

	async.NewRegister().Regist(m)
	closure.NewRegister().Regist(m)
	crypto.NewRegister().Regist(m)
	generate.NewRegister().Regist(m)
	reflection.NewRegister().Regist(m)
	sets.NewRegister().Regist(m)
	xdgspec.NewRegister().Regist(m)
}

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
	advancedExampleRegister struct{}
)

// NewRegister は、advanced パッケージ用の lib.Register を返します.
func NewRegister() interfaces.Register {
	r := new(advancedExampleRegister)
	return r
}

// Regist は、advanced パッケージ配下に存在するサンプルを登録します.
func (r *advancedExampleRegister) Regist(m interfaces.ExampleMapping) {

	register := async.NewRegister()
	register.Regist(m)

	register = closure.NewRegister()
	register.Regist(m)

	register = crypto.NewRegister()
	register.Regist(m)

	register = generate.NewRegister()
	register.Regist(m)

	register = reflection.NewRegister()
	register.Regist(m)

	register = sets.NewRegister()
	register.Regist(m)

	register = xdgspec.NewRegister()
	register.Regist(m)
}

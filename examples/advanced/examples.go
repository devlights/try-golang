package advanced

import (
	"github.com/devlights/try-golang/examples/advanced/async"
	"github.com/devlights/try-golang/examples/advanced/closure"
	"github.com/devlights/try-golang/examples/advanced/crypto"
	"github.com/devlights/try-golang/examples/advanced/deepcopy"
	"github.com/devlights/try-golang/examples/advanced/errgrp"
	"github.com/devlights/try-golang/examples/advanced/generate"
	"github.com/devlights/try-golang/examples/advanced/gocmp"
	"github.com/devlights/try-golang/examples/advanced/japanese"
	"github.com/devlights/try-golang/examples/advanced/reflection"
	"github.com/devlights/try-golang/examples/advanced/sets"
	"github.com/devlights/try-golang/examples/advanced/xdgspec"
	"github.com/devlights/try-golang/mappings"
)

type (
	register struct{}
)

// NewRegister は、advanced パッケージ用の lib.Register を返します.
func NewRegister() mappings.Register {
	r := new(register)
	return r
}

// Regist は、advanced パッケージ配下に存在するサンプルを登録します.
func (r *register) Regist(m mappings.ExampleMapping) {
	async.NewRegister().Regist(m)
	closure.NewRegister().Regist(m)
	crypto.NewRegister().Regist(m)
	deepcopy.NewRegister().Regist(m)
	errgrp.NewRegister().Regist(m)
	generate.NewRegister().Regist(m)
	gocmp.NewRegister().Regist(m)
	japanese.NewRegister().Regist(m)
	reflection.NewRegister().Regist(m)
	sets.NewRegister().Regist(m)
	xdgspec.NewRegister().Regist(m)
}

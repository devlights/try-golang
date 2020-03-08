package basic

import (
	"github.com/devlights/try-golang/basic/array"
	"github.com/devlights/try-golang/basic/builtins"
	"github.com/devlights/try-golang/basic/comments"
	"github.com/devlights/try-golang/basic/constants"
	"github.com/devlights/try-golang/basic/defers"
	"github.com/devlights/try-golang/basic/enum"
	"github.com/devlights/try-golang/basic/errs"
	"github.com/devlights/try-golang/basic/fileio"
	"github.com/devlights/try-golang/basic/filepaths"
	"github.com/devlights/try-golang/basic/functions"
	"github.com/devlights/try-golang/basic/helloworld"
	"github.com/devlights/try-golang/basic/ifs"
	"github.com/devlights/try-golang/basic/imports"
	"github.com/devlights/try-golang/basic/literals"
	"github.com/devlights/try-golang/basic/logging"
	"github.com/devlights/try-golang/basic/maps"
	"github.com/devlights/try-golang/basic/maths"
	"github.com/devlights/try-golang/basic/runtimes"
	"github.com/devlights/try-golang/basic/scope"
	"github.com/devlights/try-golang/basic/slices"
	"github.com/devlights/try-golang/basic/sorts"
	"github.com/devlights/try-golang/basic/stdin"
	"github.com/devlights/try-golang/basic/stdout"
	"github.com/devlights/try-golang/basic/strconvs"
	"github.com/devlights/try-golang/basic/strs"
	"github.com/devlights/try-golang/basic/structs"
	"github.com/devlights/try-golang/basic/system"
	"github.com/devlights/try-golang/basic/times"
	"github.com/devlights/try-golang/basic/types"
	"github.com/devlights/try-golang/basic/unsafes"
	"github.com/devlights/try-golang/basic/variables"
	"github.com/devlights/try-golang/interfaces"
)

type (
	register struct{}
)

// NewRegister は、basic パッケージ用の lib.Register を返します.
func NewRegister() interfaces.Register {
	r := new(register)
	return r
}

// Regist は、basic パッケージ配下に存在するサンプルを登録します.
func (r *register) Regist(m interfaces.ExampleMapping) {

	array.NewRegister().Regist(m)
	builtins.NewRegister().Regist(m)
	comments.NewRegister().Regist(m)
	constants.NewRegister().Regist(m)
	defers.NewRegister().Regist(m)
	enum.NewRegister().Regist(m)
	errs.NewRegister().Regist(m)
	fileio.NewRegister().Regist(m)
	filepaths.NewRegister().Regist(m)
	functions.NewRegister().Regist(m)
	helloworld.NewRegister().Regist(m)
	ifs.NewRegister().Regist(m)
	imports.NewRegister().Regist(m)
	literals.NewRegister().Regist(m)
	logging.NewRegister().Regist(m)
	maps.NewRegister().Regist(m)
	maths.NewRegister().Regist(m)
	runtimes.NewRegister().Regist(m)
	scope.NewRegister().Regist(m)
	slices.NewRegister().Regist(m)
	sorts.NewRegister().Regist(m)
	stdin.NewRegister().Regist(m)
	stdout.NewRegister().Regist(m)
	strconvs.NewRegister().Regist(m)
	strs.NewRegister().Regist(m)
	structs.NewRegister().Regist(m)
	system.NewRegister().Regist(m)
	times.NewRegister().Regist(m)
	types.NewRegister().Regist(m)
	unsafes.NewRegister().Regist(m)
	variables.NewRegister().Regist(m)
}

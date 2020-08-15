package basic

import (
	"github.com/devlights/try-golang/examples/basic/array"
	"github.com/devlights/try-golang/examples/basic/binaries"
	"github.com/devlights/try-golang/examples/basic/builtins"
	"github.com/devlights/try-golang/examples/basic/byteop"
	"github.com/devlights/try-golang/examples/basic/comments"
	"github.com/devlights/try-golang/examples/basic/constants"
	"github.com/devlights/try-golang/examples/basic/convert"
	"github.com/devlights/try-golang/examples/basic/defers"
	"github.com/devlights/try-golang/examples/basic/enum"
	"github.com/devlights/try-golang/examples/basic/errs"
	"github.com/devlights/try-golang/examples/basic/fileio"
	"github.com/devlights/try-golang/examples/basic/filepaths"
	"github.com/devlights/try-golang/examples/basic/functions"
	"github.com/devlights/try-golang/examples/basic/helloworld"
	"github.com/devlights/try-golang/examples/basic/ifs"
	"github.com/devlights/try-golang/examples/basic/imports"
	"github.com/devlights/try-golang/examples/basic/literals"
	"github.com/devlights/try-golang/examples/basic/logging"
	"github.com/devlights/try-golang/examples/basic/maps"
	"github.com/devlights/try-golang/examples/basic/maths"
	"github.com/devlights/try-golang/examples/basic/network"
	"github.com/devlights/try-golang/examples/basic/runtimes"
	"github.com/devlights/try-golang/examples/basic/scope"
	"github.com/devlights/try-golang/examples/basic/slices"
	"github.com/devlights/try-golang/examples/basic/sorts"
	"github.com/devlights/try-golang/examples/basic/stdin"
	"github.com/devlights/try-golang/examples/basic/stdout"
	"github.com/devlights/try-golang/examples/basic/strconvs"
	"github.com/devlights/try-golang/examples/basic/streams"
	"github.com/devlights/try-golang/examples/basic/strs"
	"github.com/devlights/try-golang/examples/basic/structs"
	"github.com/devlights/try-golang/examples/basic/system"
	"github.com/devlights/try-golang/examples/basic/times"
	"github.com/devlights/try-golang/examples/basic/types"
	"github.com/devlights/try-golang/examples/basic/unsafes"
	"github.com/devlights/try-golang/examples/basic/variables"
	"github.com/devlights/try-golang/mappings"
)

type (
	register struct{}
)

// NewRegister は、basic パッケージ用の lib.Register を返します.
func NewRegister() mappings.Register {
	r := new(register)
	return r
}

// Regist は、basic パッケージ配下に存在するサンプルを登録します.
func (r *register) Regist(m mappings.ExampleMapping) {

	array.NewRegister().Regist(m)
	binaries.NewRegister().Regist(m)
	builtins.NewRegister().Regist(m)
	byteop.NewRegister().Regist(m)
	comments.NewRegister().Regist(m)
	constants.NewRegister().Regist(m)
	convert.NewRegister().Regist(m)
	defers.NewRegister().Regist(m)
	enum.NewRegister().Regist(m)
	errs.NewRegister().Regist(m)
	fileio.NewRegister().Regist(m)
	filepaths.NewRegister().Regist(m)
	functions.NewRegister().Regist(m)
	helloworld.NewRegister().Regist(m)
	ifs.NewRegister().Regist(m)
	imports.NewRegister().Regist(m)
	streams.NewRegister().Regist(m)
	literals.NewRegister().Regist(m)
	logging.NewRegister().Regist(m)
	maps.NewRegister().Regist(m)
	maths.NewRegister().Regist(m)
	network.NewRegister().Regist(m)
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

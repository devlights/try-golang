package basic

import (
	"github.com/devlights/try-golang/basic/array_"
	"github.com/devlights/try-golang/basic/builtin_"
	"github.com/devlights/try-golang/basic/comments"
	"github.com/devlights/try-golang/basic/constants"
	"github.com/devlights/try-golang/basic/defer_"
	"github.com/devlights/try-golang/basic/error_"
	"github.com/devlights/try-golang/basic/functions"
	"github.com/devlights/try-golang/basic/helloworld"
	"github.com/devlights/try-golang/basic/import_"
	"github.com/devlights/try-golang/basic/interface_"
	"github.com/devlights/try-golang/basic/io_"
	"github.com/devlights/try-golang/basic/iota_"
	"github.com/devlights/try-golang/basic/literals"
	"github.com/devlights/try-golang/basic/log_"
	"github.com/devlights/try-golang/basic/map_"
	"github.com/devlights/try-golang/basic/math_"
	"github.com/devlights/try-golang/basic/os_"
	"github.com/devlights/try-golang/basic/runtime_"
	"github.com/devlights/try-golang/basic/scope"
	"github.com/devlights/try-golang/basic/slice_"
	"github.com/devlights/try-golang/basic/sort_"
	"github.com/devlights/try-golang/basic/stdin"
	"github.com/devlights/try-golang/basic/stdout"
	"github.com/devlights/try-golang/basic/strconv_"
	"github.com/devlights/try-golang/basic/string_"
	"github.com/devlights/try-golang/basic/struct_"
	"github.com/devlights/try-golang/basic/time_"
	"github.com/devlights/try-golang/basic/type_"
	"github.com/devlights/try-golang/basic/unsafe_"
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

	array_.NewRegister().Regist(m)
	builtin_.NewRegister().Regist(m)
	comments.NewRegister().Regist(m)
	constants.NewRegister().Regist(m)
	defer_.NewRegister().Regist(m)
	error_.NewRegister().Regist(m)
	functions.NewRegister().Regist(m)
	helloworld.NewRegister().Regist(m)
	import_.NewRegister().Regist(m)
	interface_.NewRegister().Regist(m)
	io_.NewRegister().Regist(m)
	iota_.NewRegister().Regist(m)
	literals.NewRegister().Regist(m)
	log_.NewRegister().Regist(m)
	map_.NewRegister().Regist(m)
	math_.NewRegister().Regist(m)
	os_.NewRegister().Regist(m)
	runtime_.NewRegister().Regist(m)
	scope.NewRegister().Regist(m)
	slice_.NewRegister().Regist(m)
	sort_.NewRegister().Regist(m)
	stdin.NewRegister().Regist(m)
	stdout.NewRegister().Regist(m)
	strconv_.NewRegister().Regist(m)
	string_.NewRegister().Regist(m)
	struct_.NewRegister().Regist(m)
	time_.NewRegister().Regist(m)
	type_.NewRegister().Regist(m)
	unsafe_.NewRegister().Regist(m)
	variables.NewRegister().Regist(m)
}

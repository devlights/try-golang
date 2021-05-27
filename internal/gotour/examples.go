package gotour

import (
	"github.com/devlights/try-golang/internal/gotour/gotour01"
	"github.com/devlights/try-golang/internal/gotour/gotour02"
	"github.com/devlights/try-golang/internal/gotour/gotour03"
	"github.com/devlights/try-golang/internal/gotour/gotour04"
	"github.com/devlights/try-golang/internal/gotour/gotour05"
	"github.com/devlights/try-golang/internal/gotour/gotour06"
	"github.com/devlights/try-golang/internal/gotour/gotour07"
	"github.com/devlights/try-golang/internal/gotour/gotour08"
	"github.com/devlights/try-golang/internal/gotour/gotour09"
	"github.com/devlights/try-golang/internal/gotour/gotour10"
	"github.com/devlights/try-golang/internal/gotour/gotour11"
	"github.com/devlights/try-golang/internal/gotour/gotour12"
	"github.com/devlights/try-golang/internal/gotour/gotour13"
	"github.com/devlights/try-golang/internal/gotour/gotour14"
	"github.com/devlights/try-golang/internal/gotour/gotour15"
	"github.com/devlights/try-golang/internal/gotour/gotour16"
	"github.com/devlights/try-golang/internal/gotour/gotour17"
	"github.com/devlights/try-golang/internal/gotour/gotour18"
	"github.com/devlights/try-golang/internal/gotour/gotour19"
	"github.com/devlights/try-golang/internal/gotour/gotour20"
	"github.com/devlights/try-golang/internal/gotour/gotour21"
	"github.com/devlights/try-golang/internal/gotour/gotour22"
	"github.com/devlights/try-golang/internal/gotour/gotour23"
	"github.com/devlights/try-golang/internal/gotour/gotour24"
	"github.com/devlights/try-golang/internal/gotour/gotour25"
	"github.com/devlights/try-golang/internal/gotour/gotour26"
	"github.com/devlights/try-golang/internal/gotour/gotour27"
	"github.com/devlights/try-golang/internal/gotour/gotour28"
	"github.com/devlights/try-golang/internal/gotour/gotour29"
	"github.com/devlights/try-golang/pkg/mappings"
)

type (
	gotourExampleRegister struct{}
)

// NewRegister は、gotour パッケージ用の lib.Register を返します.
func NewRegister() mappings.Register {
	r := new(gotourExampleRegister)
	return r
}

// Regist は、gotour パッケージ配下に存在するサンプルを登録します.
func (r *gotourExampleRegister) Regist(m mappings.ExampleMapping) {
	m["gotour_helloworld"] = gotour01.HelloWorld
	m["gotour_import"] = gotour02.Import
	m["gotour_scope"] = gotour03.Scope
	m["gotour_functions"] = gotour04.Functions
	m["gotour_basictypes"] = gotour05.BasicTypes
	m["gotour_zerovalue"] = gotour06.ZeroValue
	m["gotour_typeconvert_basictypes"] = gotour07.TypeConvertBasicTypes
	m["gotour_const"] = gotour08.Constant
	m["gotour_forloop"] = gotour09.ForLoop
	m["gotour_if"] = gotour10.If
	m["gotour_switch"] = gotour11.Switch
	m["gotour_defer"] = gotour12.Defer
	m["gotour_pointer"] = gotour13.Pointer
	m["gotour_struct"] = gotour14.Struct
	m["gotour_array"] = gotour15.Array
	m["gotour_slice"] = gotour16.Slice
	m["gotour_map"] = gotour17.Map
	m["gotour_method"] = gotour18.Method
	m["gotour_interface"] = gotour19.Interface
	m["gotour_empty_interface"] = gotour20.EmptyInterface
	m["gotour_type_assertion"] = gotour21.TypeAssertion
	m["gotour_type_switch"] = gotour22.TypeSwitch
	m["gotour_stringer"] = gotour23.Stringer
	m["gotour_error"] = gotour24.Error
	m["gotour_reader"] = gotour25.Reader
	m["gotour_goroutine"] = gotour26.Goroutine
	m["gotour_channels"] = gotour27.Channels
	m["gotour_select"] = gotour28.Select
	m["gotour_mutex"] = gotour29.Mutex
}

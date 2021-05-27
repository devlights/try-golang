package gotour

import (
	"github.com/devlights/try-golang/examples/gotour/gotour01"
	"github.com/devlights/try-golang/examples/gotour/gotour02"
	"github.com/devlights/try-golang/examples/gotour/gotour03"
	"github.com/devlights/try-golang/examples/gotour/gotour04"
	"github.com/devlights/try-golang/examples/gotour/gotour05"
	"github.com/devlights/try-golang/examples/gotour/gotour06"
	"github.com/devlights/try-golang/examples/gotour/gotour07"
	"github.com/devlights/try-golang/examples/gotour/gotour08"
	"github.com/devlights/try-golang/examples/gotour/gotour09"
	"github.com/devlights/try-golang/examples/gotour/gotour10"
	"github.com/devlights/try-golang/examples/gotour/gotour11"
	"github.com/devlights/try-golang/examples/gotour/gotour12"
	"github.com/devlights/try-golang/examples/gotour/gotour13"
	"github.com/devlights/try-golang/examples/gotour/gotour14"
	"github.com/devlights/try-golang/examples/gotour/gotour15"
	"github.com/devlights/try-golang/examples/gotour/gotour16"
	"github.com/devlights/try-golang/examples/gotour/gotour17"
	"github.com/devlights/try-golang/examples/gotour/gotour18"
	"github.com/devlights/try-golang/examples/gotour/gotour19"
	"github.com/devlights/try-golang/examples/gotour/gotour20"
	"github.com/devlights/try-golang/examples/gotour/gotour21"
	"github.com/devlights/try-golang/examples/gotour/gotour22"
	"github.com/devlights/try-golang/examples/gotour/gotour23"
	"github.com/devlights/try-golang/examples/gotour/gotour24"
	"github.com/devlights/try-golang/examples/gotour/gotour25"
	"github.com/devlights/try-golang/examples/gotour/gotour26"
	"github.com/devlights/try-golang/examples/gotour/gotour27"
	"github.com/devlights/try-golang/examples/gotour/gotour28"
	"github.com/devlights/try-golang/examples/gotour/gotour29"
	"github.com/devlights/try-golang/mappings"
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
	m["gotour_01_helloworld"] = gotour01.HelloWorld
	m["gotour_02_import"] = gotour02.Import
	m["gotour_03_scope"] = gotour03.Scope
	m["gotour_04_functions"] = gotour04.Functions
	m["gotour_05_basictypes"] = gotour05.BasicTypes
	m["gotour_06_zerovalue"] = gotour06.ZeroValue
	m["gotour_07_typeconvert_basictypes"] = gotour07.TypeConvertBasicTypes
	m["gotour_08_const"] = gotour08.Constant
	m["gotour_09_forloop"] = gotour09.ForLoop
	m["gotour_10_if"] = gotour10.If
	m["gotour_11_switch"] = gotour11.Switch
	m["gotour_12_defer"] = gotour12.Defer
	m["gotour_13_pointer"] = gotour13.Pointer
	m["gotour_14_struct"] = gotour14.Struct
	m["gotour_15_array"] = gotour15.Array
	m["gotour_16_slice"] = gotour16.Slice
	m["gotour_17_map"] = gotour17.Map
	m["gotour_18_method"] = gotour18.Method
	m["gotour_19_interface"] = gotour19.Interface
	m["gotour_20_empty_interface"] = gotour20.EmptyInterface
	m["gotour_21_type_assertion"] = gotour21.TypeAssertion
	m["gotour_22_type_switch"] = gotour22.TypeSwitch
	m["gotour_23_stringer"] = gotour23.Stringer
	m["gotour_24_error"] = gotour24.Error
	m["gotour_25_reader"] = gotour25.Reader
	m["gotour_26_goroutine"] = gotour26.Goroutine
	m["gotour_27_channels"] = gotour27.Channels
	m["gotour_28_select"] = gotour28.Select
	m["gotour_29_mutex"] = gotour29.Mutex
}

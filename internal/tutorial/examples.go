package tutorial

import (
	"github.com/devlights/try-golang/mappings"
	"github.com/devlights/try-golang/tutorial/tutorial01"
	"github.com/devlights/try-golang/tutorial/tutorial02"
	"github.com/devlights/try-golang/tutorial/tutorial03"
	"github.com/devlights/try-golang/tutorial/tutorial04"
	"github.com/devlights/try-golang/tutorial/tutorial05"
	"github.com/devlights/try-golang/tutorial/tutorial06"
	"github.com/devlights/try-golang/tutorial/tutorial07"
	"github.com/devlights/try-golang/tutorial/tutorial08"
	"github.com/devlights/try-golang/tutorial/tutorial09"
	"github.com/devlights/try-golang/tutorial/tutorial10"
	"github.com/devlights/try-golang/tutorial/tutorial11"
	"github.com/devlights/try-golang/tutorial/tutorial12"
	"github.com/devlights/try-golang/tutorial/tutorial13"
	"github.com/devlights/try-golang/tutorial/tutorial14"
	"github.com/devlights/try-golang/tutorial/tutorial15"
	"github.com/devlights/try-golang/tutorial/tutorial16"
	"github.com/devlights/try-golang/tutorial/tutorial17"
	"github.com/devlights/try-golang/tutorial/tutorial18"
	"github.com/devlights/try-golang/tutorial/tutorial19"
	"github.com/devlights/try-golang/tutorial/tutorial20"
	"github.com/devlights/try-golang/tutorial/tutorial21"
	"github.com/devlights/try-golang/tutorial/tutorial22"
	"github.com/devlights/try-golang/tutorial/tutorial23"
	"github.com/devlights/try-golang/tutorial/tutorial24"
	"github.com/devlights/try-golang/tutorial/tutorial25"
	"github.com/devlights/try-golang/tutorial/tutorial26"
	"github.com/devlights/try-golang/tutorial/tutorial27"
	"github.com/devlights/try-golang/tutorial/tutorial28"
	"github.com/devlights/try-golang/tutorial/tutorial29"
)

type (
	tutorialExampleRegister struct{}
)

// NewRegister は、tutorial パッケージ用の lib.Register を返します.
func NewRegister() mappings.Register {
	r := new(tutorialExampleRegister)
	return r
}

// Regist は、tutorial パッケージ配下に存在するサンプルを登録します.
func (r *tutorialExampleRegister) Regist(m mappings.ExampleMapping) {
	m["tutorial_gotour_helloworld"] = tutorial01.HelloWorld
	m["tutorial_gotour_import"] = tutorial02.Import
	m["tutorial_gotour_scope"] = tutorial03.Scope
	m["tutorial_gotour_functions"] = tutorial04.Functions
	m["tutorial_gotour_basictypes"] = tutorial05.BasicTypes
	m["tutorial_gotour_zerovalue"] = tutorial06.ZeroValue
	m["tutorial_gotour_typeconvert_basictypes"] = tutorial07.TypeConvertBasicTypes
	m["tutorial_gotour_const"] = tutorial08.Constant
	m["tutorial_gotour_forloop"] = tutorial09.ForLoop
	m["tutorial_gotour_if"] = tutorial10.If
	m["tutorial_gotour_switch"] = tutorial11.Switch
	m["tutorial_gotour_defer"] = tutorial12.Defer
	m["tutorial_gotour_pointer"] = tutorial13.Pointer
	m["tutorial_gotour_struct"] = tutorial14.Struct
	m["tutorial_gotour_array"] = tutorial15.Array
	m["tutorial_gotour_slice"] = tutorial16.Slice
	m["tutorial_gotour_map"] = tutorial17.Map
	m["tutorial_gotour_method"] = tutorial18.Method
	m["tutorial_gotour_interface"] = tutorial19.Interface
	m["tutorial_gotour_empty_interface"] = tutorial20.EmptyInterface
	m["tutorial_gotour_type_assertion"] = tutorial21.TypeAssertion
	m["tutorial_gotour_type_switch"] = tutorial22.TypeSwitch
	m["tutorial_gotour_stringer"] = tutorial23.Stringer
	m["tutorial_gotour_error"] = tutorial24.Error
	m["tutorial_gotour_reader"] = tutorial25.Reader
	m["tutorial_gotour_goroutine"] = tutorial26.Goroutine
	m["tutorial_gotour_channels"] = tutorial27.Channels
	m["tutorial_gotour_select"] = tutorial28.Select
	m["tutorial_gotour_mutex"] = tutorial29.Mutex
}

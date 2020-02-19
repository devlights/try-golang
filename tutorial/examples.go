package tutorial

import (
	"github.com/devlights/try-golang/interfaces"
)

type (
	tutorialExampleRegister struct{}
)

// NewRegister は、tutorial パッケージ用の lib.Register を返します.
func NewRegister() interfaces.Register {
	r := new(tutorialExampleRegister)
	return r
}

func (r *tutorialExampleRegister) Regist(m interfaces.ExampleMapping) {
	m["tutorial_gotour_helloworld"] = HelloWorld
	m["tutorial_gotour_import"] = Import
	m["tutorial_gotour_scope"] = Scope
	m["tutorial_gotour_functions"] = Functions
	m["tutorial_gotour_basictypes"] = BasicTypes
	m["tutorial_gotour_zerovalue"] = ZeroValue
	m["tutorial_gotour_typeconvert_basictypes"] = TypeConvertBasicTypes
	m["tutorial_gotour_const"] = Constant
	m["tutorial_gotour_forloop"] = ForLoop
	m["tutorial_gotour_if"] = If
	m["tutorial_gotour_switch"] = Switch
	m["tutorial_gotour_defer"] = Defer
	m["tutorial_gotour_pointer"] = Pointer
	m["tutorial_gotour_struct"] = Struct
	m["tutorial_gotour_array"] = Array
	m["tutorial_gotour_slice"] = Slice
	m["tutorial_gotour_map"] = Map
	m["tutorial_gotour_method"] = Method
	m["tutorial_gotour_interface"] = Interface
	m["tutorial_gotour_empty_interface"] = EmptyInterface
	m["tutorial_gotour_type_assertion"] = TypeAssertion
	m["tutorial_gotour_type_switch"] = TypeSwitch
	m["tutorial_gotour_stringer"] = Stringer
	m["tutorial_gotour_error"] = Error
	m["tutorial_gotour_reader"] = Reader
	m["tutorial_gotour_goroutine"] = Goroutine
	m["tutorial_gotour_channels"] = Channels
	m["tutorial_gotour_select"] = Select
	m["tutorial_gotour_mutex"] = Mutex
}

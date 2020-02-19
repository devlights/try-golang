package effectivego

import (
	"github.com/devlights/try-golang/interfaces"
)

type (
	effectivegoExampleRegister struct{}
)

// NewRegister は、effectivego パッケージ用の lib.Register を返します.
func NewRegister() interfaces.Register {
	r := new(effectivegoExampleRegister)
	return r
}

func (r *effectivegoExampleRegister) Regist(m interfaces.ExampleMapping) {
	m["effective_go_intro"] = Introduction
	m["effective_go_formatting"] = Formatting
	m["effective_go_comment"] = Commentary
	m["effective_go_names"] = Names
	m["effective_go_semicolon"] = Semicolons
	m["effective_go_control"] = ControlStructure
	m["effective_go_functions"] = Functions
	m["effective_go_defer"] = Defer
	m["effective_go_allocation_with_new"] = AllocationWithNew
	m["effective_go_constructors"] = Constructors
	m["effective_go_allocation_with_make"] = AllocationWithMake
	m["effective_go_arrays"] = Arrays
	m["effective_go_slices"] = Slices
	m["effective_go_two_dimentional_slices"] = TwoDimentionalSlices
	m["effective_go_maps"] = Maps
	m["effective_go_printing"] = Printing
	m["effective_go_append"] = Append
	m["effective_go_constants"] = Constants
	m["effective_go_methods"] = Methods
}

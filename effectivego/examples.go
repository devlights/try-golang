package effectivego

import (
	"github.com/devlights/try-golang/effectivego/effectivego01"
	"github.com/devlights/try-golang/effectivego/effectivego02"
	"github.com/devlights/try-golang/effectivego/effectivego03"
	"github.com/devlights/try-golang/effectivego/effectivego04"
	"github.com/devlights/try-golang/effectivego/effectivego05"
	"github.com/devlights/try-golang/effectivego/effectivego06"
	"github.com/devlights/try-golang/effectivego/effectivego07"
	"github.com/devlights/try-golang/effectivego/effectivego08"
	"github.com/devlights/try-golang/effectivego/effectivego09"
	"github.com/devlights/try-golang/effectivego/effectivego10"
	"github.com/devlights/try-golang/effectivego/effectivego11"
	"github.com/devlights/try-golang/effectivego/effectivego12"
	"github.com/devlights/try-golang/effectivego/effectivego13"
	"github.com/devlights/try-golang/effectivego/effectivego14"
	"github.com/devlights/try-golang/effectivego/effectivego15"
	"github.com/devlights/try-golang/effectivego/effectivego16"
	"github.com/devlights/try-golang/effectivego/effectivego17"
	"github.com/devlights/try-golang/effectivego/effectivego18"
	"github.com/devlights/try-golang/effectivego/effectivego19"
	"github.com/devlights/try-golang/effectivego/effectivego20"
	"github.com/devlights/try-golang/mappings"
)

type (
	effectivegoExampleRegister struct{}
)

// NewRegister は、effectivego パッケージ用の lib.Register を返します.
func NewRegister() mappings.Register {
	r := new(effectivegoExampleRegister)
	return r
}

// Regist は、effectivego パッケージ配下に存在するサンプルを登録します.
func (r *effectivegoExampleRegister) Regist(m mappings.ExampleMapping) {
	m["effective_go_intro"] = effectivego01.Introduction
	m["effective_go_formatting"] = effectivego02.Formatting
	m["effective_go_comment"] = effectivego03.Commentary
	m["effective_go_names"] = effectivego04.Names
	m["effective_go_semicolon"] = effectivego05.Semicolons
	m["effective_go_control"] = effectivego06.ControlStructure
	m["effective_go_functions"] = effectivego07.Functions
	m["effective_go_defer"] = effectivego08.Defer
	m["effective_go_allocation_with_new"] = effectivego09.AllocationWithNew
	m["effective_go_constructors"] = effectivego10.Constructors
	m["effective_go_allocation_with_make"] = effectivego11.AllocationWithMake
	m["effective_go_arrays"] = effectivego12.Arrays
	m["effective_go_slices"] = effectivego13.Slices
	m["effective_go_two_dimensional_slices"] = effectivego14.TwoDimensionalSlices
	m["effective_go_maps"] = effectivego15.Maps
	m["effective_go_printing"] = effectivego16.Printing
	m["effective_go_append"] = effectivego17.Append
	m["effective_go_constants"] = effectivego18.Constants
	m["effective_go_methods"] = effectivego19.Methods
	m["effective_go_interfaces"] = effectivego20.Interfaces
}

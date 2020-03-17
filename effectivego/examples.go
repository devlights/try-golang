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
	"github.com/devlights/try-golang/effectivego/effectivego21"
	"github.com/devlights/try-golang/effectivego/effectivego22"
	"github.com/devlights/try-golang/effectivego/effectivego23"
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
	m["effectivego_01_intro"] = effectivego01.Introduction
	m["effectivego_02_formatting"] = effectivego02.Formatting
	m["effectivego_03_comment"] = effectivego03.Commentary
	m["effectivego_04_names"] = effectivego04.Names
	m["effectivego_05_semicolon"] = effectivego05.Semicolons
	m["effectivego_06_control"] = effectivego06.ControlStructure
	m["effectivego_07_functions"] = effectivego07.Functions
	m["effectivego_08_defer"] = effectivego08.Defer
	m["effectivego_09_allocation_with_new"] = effectivego09.AllocationWithNew
	m["effectivego_10_constructors"] = effectivego10.Constructors
	m["effectivego_11_allocation_with_make"] = effectivego11.AllocationWithMake
	m["effectivego_12_arrays"] = effectivego12.Arrays
	m["effectivego_13_slices"] = effectivego13.Slices
	m["effectivego_14_two_dimensional_slices"] = effectivego14.TwoDimensionalSlices
	m["effectivego_15_maps"] = effectivego15.Maps
	m["effectivego_16_printing"] = effectivego16.Printing
	m["effectivego_17_append"] = effectivego17.Append
	m["effectivego_18_constants"] = effectivego18.Constants
	m["effectivego_19_methods"] = effectivego19.Methods
	m["effectivego_20_interfaces"] = effectivego20.Interfaces
	m["effectivego_21_interface_conversion"] = effectivego21.InterfaceConversion
	m["effectivego_22_generality"] = effectivego22.Generality
	m["effectivego_23_interface_check"] = effectivego23.InterfaceCheck
}

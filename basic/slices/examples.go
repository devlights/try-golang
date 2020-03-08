package slices

import (
	"github.com/devlights/try-golang/mappings"
)

type (
	register struct{}
)

func NewRegister() mappings.Register {
	return new(register)
}

func (r *register) Regist(m mappings.ExampleMapping) {
	m["slice_basic01"] = Basic01
	m["slice_basic02"] = Basic02
	m["slice_basic03"] = Basic03
	m["slice_basic04"] = Basic04
	m["slice_basic05"] = Basic05
	m["slice_reverse"] = SliceReverse
	m["slice_append"] = SliceAppend
	m["slice_pointer"] = SlicePointer
	m["slice_copy"] = SliceCopy
	m["slice_clear"] = SliceClear
	m["slice_deep_equal"] = SliceDeepEqual
}

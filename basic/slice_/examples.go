package slice_

import (
	"github.com/devlights/try-golang/interfaces"
)

type (
	register struct{}
)

func NewRegister() interfaces.Register {
	return new(register)
}

func (r *register) Regist(m interfaces.ExampleMapping) {
	m["slice01"] = Slice01
	m["slice02"] = Slice02
	m["slice03"] = Slice03
	m["slice04"] = Slice04
	m["slice05"] = Slice05
	m["slice_reverse"] = SliceReverse
	m["slice_append"] = SliceAppend
	m["slice_pointer"] = SlicePointer
	m["slice_copy"] = SliceCopy
	m["slice_clear"] = SliceClear
	m["slice_deep_equal"] = SliceDeepEqual
}

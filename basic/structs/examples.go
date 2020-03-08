package structs

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
	m["struct_basic01"] = Basic01
	m["struct_basic02"] = Basic02
	m["struct_basic03"] = Basic03
	m["struct_basic04"] = Basic04
	m["struct_anonymous_struct"] = StructAnonymousStruct
	m["struct_empty_struct"] = EmptyStruct
	m["struct_deep_equal"] = StructDeepEqual
}

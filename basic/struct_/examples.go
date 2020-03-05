package struct_

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
	m["struct01"] = Struct01
	m["struct02"] = Struct02
	m["struct03"] = Struct03
	m["struct04"] = Struct04
	m["struct_anonymous_struct"] = StructAnonymousStruct
	m["struct_empty_struct"] = EmptyStruct
	m["struct_deep_equal"] = StructDeepEqual
}

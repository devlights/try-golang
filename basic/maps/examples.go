package maps

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
	m["map_basic"] = MapBasic
	m["map_for"] = MapFor
	m["map_initialize"] = MapInitialize
	m["map_delete"] = MapDelete
	m["map_access"] = MapAccess
	m["map_deep_equal"] = MapDeepEqual
}

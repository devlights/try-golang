package sets

import (
	"github.com/devlights/try-golang/mappings"
)

type (
	register struct{}
)

func NewRegister() mappings.Register {
	r := new(register)
	return r
}

func (r *register) Regist(m mappings.ExampleMapping) {
	m["set01"] = Set01
	m["set02"] = Set02
	m["set03"] = Set03
	m["set04"] = Set04
	m["set05"] = Set05
	m["mapset_01"] = MapSet01
}

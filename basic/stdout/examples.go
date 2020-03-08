package stdout

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
	m["printf01"] = Printf01
	m["printf02"] = Printf02
	m["printf03"] = Printf03
	m["println01"] = Println01
}

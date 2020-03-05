package stdout

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
	m["printf01"] = Printf01
	m["printf02"] = Printf02
	m["printf03"] = Printf03
	m["println01"] = Println01
}

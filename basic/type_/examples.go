package type_

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
	m["type01"] = Type01
}

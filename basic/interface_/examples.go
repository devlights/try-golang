package interface_

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
	m["interface_basic"] = Basic
	m["interface_composition"] = Composition
	m["interface_ducktyping"] = DuckTyping
}

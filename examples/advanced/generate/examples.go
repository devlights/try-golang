package generate

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
	m["generate_generic_stack"] = UseGenericStack
	m["generate_generic_queue"] = UseGenericQueue
}

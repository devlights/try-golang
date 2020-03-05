package generate

import "github.com/devlights/try-golang/interfaces"

type (
	register struct{}
)

func NewRegister() interfaces.Register {
	r := new(register)
	return r
}

func (r *register) Regist(m interfaces.ExampleMapping) {
	m["generate_generic_stack"] = UseGenericStack
	m["generate_generic_queue"] = UseGenericQueue
}

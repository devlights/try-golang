package defers

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
	m["defer_basic_usage"] = Basic
	m["defer_in_loop"] = DeferInLoop
	m["defer_in_loop_manyfiles"] = DeferInLoopManyFiles
}

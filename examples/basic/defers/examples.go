package defers

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
	m["defer_basic_usage"] = Basic
	m["defer_in_loop"] = DeferInLoop
	m["defer_in_loop_manyfiles"] = DeferInLoopManyFiles
}

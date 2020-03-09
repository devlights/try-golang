package runtimes

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
	m["runtime_version"] = RuntimeVersion
	m["runtime_memorystats"] = RuntimeMemoryStats
}

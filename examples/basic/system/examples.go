package system

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
	m["os_getpagesize"] = PageSize
}
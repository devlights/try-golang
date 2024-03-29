package tsvop

import "github.com/devlights/try-golang/mapping"

type (
	register struct{}
)

var (
	_ mapping.Register = (*register)(nil)
)

func NewRegister() mapping.Register {
	return new(register)
}

func (*register) Regist(m mapping.ExampleMapping) {
	m["tsvop_read"] = Read
	m["tsvop_write"] = Write
}

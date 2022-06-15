package sjis

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
	m["fileio_japanese_sjis_read"] = Read
	m["fileio_japanese_sjis_write"] = Write
}

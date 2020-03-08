package fileio

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
	m["fileio_open_read"] = OpenRead
	m["fileio_open_write"] = OpenWrite
	m["fileio_stat_mkdir_removeall"] = StatMkdirRemoveAll
	m["fileio_stat"] = Stat
}

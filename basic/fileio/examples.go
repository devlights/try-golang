package fileio

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
	m["fileio_open_read"] = OpenRead
	m["fileio_open_write"] = OpenWrite
	m["fileio_stat_mkdir_removeall"] = StatMkdirRemoveAll
	m["fileio_stat"] = Stat
}

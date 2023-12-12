package stat

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
	m["fileio_stat_mkdir_removeall"] = MkdirRemoveAll
	m["fileio_stat"] = Stat
	m["fileio_stat_permission"] = Permission
}

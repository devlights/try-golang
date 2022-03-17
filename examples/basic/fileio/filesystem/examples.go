package filesystem

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
	m["fileio_filesystem_dirfs"] = DirFS
	m["fileio_filesystem_listdir"] = Listdir
	m["fileio_filesystem_readdir"] = ReadDir
	m["fileio_filesystem_notexists"] = NotExists
}

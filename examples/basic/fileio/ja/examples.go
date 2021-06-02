package ja

import (
	"github.com/devlights/try-golang/examples/basic/fileio/ja/eucjp"
	"github.com/devlights/try-golang/examples/basic/fileio/ja/sjis"
	"github.com/devlights/try-golang/mapping"
)

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
	sjis.NewRegister().Regist(m)
	eucjp.NewRegister().Regist(m)
}
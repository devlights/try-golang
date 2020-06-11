package fileio

import (
	"github.com/devlights/try-golang/mappings"
)

type (
	register struct{}
)

// NewRegister -- このパッケージ用のサンプルを登録する mappings.Register を生成します。
func NewRegister() mappings.Register {
	return new(register)
}

// Regist -- 登録します.
func (r *register) Regist(m mappings.ExampleMapping) {
	m["fileio_open_read"] = OpenRead
	m["fileio_open_write"] = OpenWrite
	m["fileio_stat_mkdir_removeall"] = StatMkdirRemoveAll
	m["fileio_stat"] = Stat
}

package fileio

import "github.com/devlights/try-golang/mapping"

type (
	register struct{}
)

// NewRegister -- このパッケージ用のサンプルを登録する mapping.Register を生成します。
func NewRegister() mapping.Register {
	return new(register)
}

// Regist -- 登録します.
func (r *register) Regist(m mapping.ExampleMapping) {
	m["fileio_open_read"] = OpenRead
	m["fileio_open_read2"] = OpenRead2
	m["fileio_open_write"] = OpenWrite
	m["fileio_stat_mkdir_removeall"] = StatMkdirRemoveAll
	m["fileio_stat"] = Stat
	m["fileio_null_writer"] = NullWriter
	m["fileio_stdin_scanner"] = StdinWithScanner
	m["fileio_stat_permission"] = StatPermission
	m["fileio_fs_dirfs"] = FsDirFS
}

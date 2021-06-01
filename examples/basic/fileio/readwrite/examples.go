package readwrite

import "github.com/devlights/try-golang/mapping"

type (
	register struct{}
)

var (
	_ mapping.Register = (*register)(nil)
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
	m["fileio_null_writer"] = NullWriter
}
package sync_and_close

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
	m["fileio_sync_and_close"] = SyncAndClose
}

package ioop

import (
	"github.com/devlights/try-golang/mapping"
)

type (
	register struct{}
)

// NewRegister -- このパッケージ用のサンプルを登録する mapping.Register を生成します。
func NewRegister() mapping.Register {
	return new(register)
}

// Regist -- 登録します.
func (r *register) Regist(m mapping.ExampleMapping) {
	m["ioop_limit_read"] = LimitRead
	m["ioop_onebyte_read"] = OneByteRead
	m["ioop_gzip_and_crc"] = GzipAndCrc
}

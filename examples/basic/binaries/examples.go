package binaries

import (
	"github.com/devlights/try-golang/examples/basic/binaries/readwrite"
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
	m["binary_byteorder"] = ByteOrder
	m["binary_readwrite"] = readwrite.ReadWrite
	m["binary_using_hex_dumper"] = UsingHexDumper
}

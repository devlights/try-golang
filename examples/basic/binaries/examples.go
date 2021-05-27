package binaries

import (
	"github.com/devlights/try-golang/examples/basic/binaries/readwrite"
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
	m["binary_byteorder"] = ByteOrder
	m["binary_readwrite"] = readwrite.ReadWrite
	m["binary_using_hex_dumper"] = UsingHexDumper
}

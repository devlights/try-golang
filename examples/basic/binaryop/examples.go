package binaryop

import (
	"github.com/devlights/try-golang/examples/basic/binaryop/readwrite"
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
	m["binaryop_read"] = Read
	m["binaryop_write"] = Write
	m["binaryop_mapping"] = MapStruct
	m["binaryop_byteorder"] = ByteOrder
	m["binaryop_readwrite"] = readwrite.ReadWrite
}

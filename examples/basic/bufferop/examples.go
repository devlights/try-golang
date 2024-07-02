package bufferop

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
	m["bufferop_zero_value"] = ZeroValue
	m["bufferop_from_string"] = FromString
	m["bufferop_from_bytes"] = FromBytes
	m["bufferop_use_as_reader"] = UseAsReader
	m["bufferop_use_as_writer"] = UseAsWriter
	m["bufferop_available_buffer"] = AvailableBuffer
}

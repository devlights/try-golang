package strs

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
	m["string_rune_rawstring"] = StringRuneRawString
	m["string_to_runeslice"] = StringToRuneSlice
	m["string_rune_byte_convert"] = StringRuneByteConvert
	m["string_chop_newline"] = ChopNewLine
}

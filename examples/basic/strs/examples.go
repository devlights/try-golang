package strs

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
	m["string_rune_rawstring"] = StringRuneRawString
	m["string_to_runeslice"] = StringToRuneSlice
	m["string_rune_byte_convert"] = StringRuneByteConvert
	m["string_chop_newline"] = ChopNewLine
	m["string_using_builder"] = UsingBuilder
}

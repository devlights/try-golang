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
	m["string_rune_rawstring"] = RuneRawString
	m["string_to_runeslice"] = ToRuneSlice
	m["string_rune_byte_convert"] = RuneByteConvert
	m["string_chop_newline"] = ChopNewLine
	m["string_using_builder"] = UsingBuilder
}

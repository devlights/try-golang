package strs

import (
	"github.com/devlights/try-golang/mappings"
)

type (
	register struct{}
)

func NewRegister() mappings.Register {
	return new(register)
}

func (r *register) Regist(m mappings.ExampleMapping) {
	m["string_rune_rawstring"] = StringRuneRawString
	m["string_to_runeslice"] = StringToRuneSlice
	m["string_rune_byte_convert"] = StringRuneByteConvert
}

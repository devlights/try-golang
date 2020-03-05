package string_

import (
	"github.com/devlights/try-golang/interfaces"
)

type (
	register struct{}
)

func NewRegister() interfaces.Register {
	return new(register)
}

func (r *register) Regist(m interfaces.ExampleMapping) {
	m["string_rune_rawstring"] = StringRuneRawString
	m["string_to_runeslice"] = StringToRuneSlice
	m["string_rune_byte_convert"] = StringRuneByteConvert
}

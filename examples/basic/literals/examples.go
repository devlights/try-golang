package literals

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
	m["binary_int_literals"] = BinaryIntLiterals
	m["octal_int_literals"] = OctalIntLiterals
	m["hex_int_literals"] = HexIntLiterals
	m["digit_separator"] = DigitSeparators
}
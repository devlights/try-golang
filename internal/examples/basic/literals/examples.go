package literals

import (
	"github.com/devlights/try-golang/pkg/mappings"
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
	m["binary_int_literals"] = BinaryIntLiterals
	m["octal_int_literals"] = OctalIntLiterals
	m["hex_int_literals"] = HexIntLiterals
	m["digit_separator"] = DigitSeparators
}

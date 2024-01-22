package literals

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
	m["literals_binary_int_literals"] = BinaryIntLiterals
	m["literals_octal_int_literals"] = OctalIntLiterals
	m["literals_hex_int_literals"] = HexIntLiterals
	m["literals_digit_separator"] = DigitSeparators
}

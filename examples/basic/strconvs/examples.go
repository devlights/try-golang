package strconvs

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
	m["strconvs_hex_to_decimal"] = HexToDecimal
	m["strconvs_bin_to_decimal"] = BinToDecimal
	m["strconvs_parseint_tips_bitsize"] = ParseIntTipsBitSize
	m["strconvs_parseint_tips_base"] = ParseIntTipsBaseValue
}

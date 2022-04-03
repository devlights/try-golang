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
	m["strconvs_hex_to_dec"] = HexToDec
	m["strconvs_bin_to_dec"] = BinToDec
	m["strconvs_hex_to_bin"] = HexToBin
	m["strconvs_bin_to_hex"] = BinToHex
	m["strconvs_dec_to_dec"] = DecToDec
	m["strconvs_parseint_tips_bitsize"] = ParseIntTipsBitSize
	m["strconvs_parseint_tips_base"] = ParseIntTipsBaseValue
}

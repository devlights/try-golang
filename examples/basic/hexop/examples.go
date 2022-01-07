package hexop

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
	m["hexop_encode"] = Encode
	m["hexop_decode"] = Decode
	m["hexop_encoder"] = Encoder
	m["hexop_decoder"] = Decoder
}

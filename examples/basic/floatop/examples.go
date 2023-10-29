package floatop

import "github.com/devlights/try-golang/mapping"

type (
	register struct{}
)

// NewRegister -- このパッケージ用のサンプルを登録する mapping.Register を生成します。
func NewRegister() mapping.Register {
	return &register{}
}

// Regist -- 登録します.
func (r *register) Regist(m mapping.ExampleMapping) {
	m["floatop_order_of_computation"] = OrderOfComputation
	m["floatop_rounding_error"] = RoundingError
}

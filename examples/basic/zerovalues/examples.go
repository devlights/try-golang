package zerovalues

import (
	"github.com/devlights/try-golang/mapping"
)

type (
	register struct{}
)

// NewRegister -- このパッケージ用のサンプルを登録する mapping.Register を生成します。
func NewRegister() mapping.Register {
	return new(register)
}

// Regist -- 登録します.
func (r *register) Regist(m mapping.ExampleMapping) {
	m["zerovalues_int"] = Int
	m["zerovalues_float"] = Float
	m["zerovalues_bool"] = Bool
	m["zerovalues_string"] = String
	m["zerovalues_pointer"] = Pointer
	m["zerovalues_slice"] = Slice
	m["zerovalues_map"] = Map
}

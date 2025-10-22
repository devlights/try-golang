package unsafes

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
	m["unsafe_sizeof"] = Sizeof
	m["unsafe_string"] = UnsafeString
	m["unsafe_stringdata"] = UnsafeStringData
	m["unsafe_pointer_cast"] = PointerCast
}

package array

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
	m["array_basic_usage"] = Basic
	m["array_copy_from_slice"] = CopyFromSlice
	m["array_ellipses"] = Ellipses
	m["array_multi_demention"] = MultiDemension
}

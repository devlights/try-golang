package sorts

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
	m["sort_interface"] = SortInterface
	m["sort_slice_unstable"] = SortSliceUnStable
	m["sort_slice_stable"] = SortSliceStable
}

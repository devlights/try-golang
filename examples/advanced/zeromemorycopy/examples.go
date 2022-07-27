package zeromemorycopy

import "github.com/devlights/try-golang/mapping"

type (
	register struct{}
)

// NewRegister -- このパッケージ用のサンプルを登録する mapping.Register を生成します。
func NewRegister() mapping.Register {
	r := new(register)
	return r
}

// Regist -- 登録します.
func (r *register) Regist(m mapping.ExampleMapping) {
	m["zeromemorycopy_string_to_byteslice"] = StringToByteSlice
	m["zeromemorycopy_byteslice_to_string"] = ByteSliceToString
}

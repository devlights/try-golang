package convert

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
	m["convert_string_slice_to_interface_slice"] = StringSliceToInterfaceSlice
	m["convert_int_to_str"] = IntToStr
	m["convert_struct_to_str"] = StructToStr
}

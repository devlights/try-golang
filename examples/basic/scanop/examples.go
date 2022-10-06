package scanop

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
	m["scanop_read_one_input"] = ReadOneInput
	m["scanop_read_multi_input"] = ReadMultipleInput
	m["scanop_read_formatted_input"] = ReadFormattedInput
}

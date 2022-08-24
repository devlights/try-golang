package formatting

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
	m["formatting_adverb_asterisk"] = AdverbAsterisk
	m["formatting_adverb_explicit_argument_indexes"] = AdverbExplicitArgumentIndexes
	m["formatting_using_v"] = UsingV
	m["formatting_append"] = Append
	m["formatting_appendf"] = AppendF
	m["formatting_appendln"] = AppendLn
}

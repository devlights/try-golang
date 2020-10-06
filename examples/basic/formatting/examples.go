package formatting

import (
	"github.com/devlights/try-golang/mappings"
)

type (
	register struct{}
)

// NewRegister -- このパッケージ用のサンプルを登録する mappings.Register を生成します。
func NewRegister() mappings.Register {
	return new(register)
}

// Regist -- 登録します.
func (r *register) Regist(m mappings.ExampleMapping) {
	m["formatting_adverb_asterisk"] = AdverbAsterisk
	m["formatting_adverb_explicit_argument_indexes"] = AdverbExplicitArgumentIndexes
}

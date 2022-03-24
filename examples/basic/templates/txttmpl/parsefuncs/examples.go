package parsefuncs

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
	m["templates_parsefuncs_parsefs"] = ParseFS
	m["templates_parsefuncs_parsefiles"] = ParseFiles
	m["templates_parsefuncs_parseglob"] = ParseGlob
}
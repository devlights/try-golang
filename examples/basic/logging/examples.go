package logging

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
	m["logging_flags"] = Flags
	m["logging_prefix"] = Prefix
	m["logging_output"] = Output
	m["logging_new"] = NewLogger
	m["logging_msgprefix"] = Msgprefix
}

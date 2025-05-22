package flags

import "github.com/devlights/try-golang/mapping"

type (
	register struct{}
)

// NewRegister -- このパッケージ用のサンプルを登録する mapping.Register を生成します。
func NewRegister() mapping.Register {
	return &register{}
}

// Regist -- 登録します.
func (r *register) Regist(m mapping.ExampleMapping) {
	m["flags_flagset"] = Flagset
	m["flags_var"] = Var
	m["flags_var2"] = Var2
	m["flags_int"] = Int
	m["flags_bool"] = Bool
	m["flags_string"] = String
	m["flags_duration"] = Duration
	m["flags_func"] = Func
	m["flags_textvar"] = TextVar
	m["flags_nargs"] = Nargs
	m["flags_subcommand"] = Subcommand
}

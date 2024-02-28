package osop

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
	m["osop_mkdir"] = Mkdir
	m["osop_list_processes"] = ListProcesses
	m["osop_environ"] = Environ
	m["osop_getenv"] = GetEnv
	m["osop_lookupenv"] = LookupEnv
	m["osop_expandenv"] = ExpandEnv
	m["osop_expand"] = Expand
	m["osop_setenv"] = Setenv
	m["osop_unsetenv"] = Unsetenv
}

package runtimes

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
	m["runtime_version"] = RuntimeVersion
	m["runtime_memorystats"] = RuntimeMemoryStats
	m["runtime_gomaxprocs"] = GoMaxProcs
	m["runtime_goexit"] = Goexit
}

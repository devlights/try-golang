package helloworld

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
	m["helloworld_sync"] = Sync
	m["helloworld_async"] = Async
	m["helloworld_async2"] = Async2
	m["helloworld_mixed"] = Mixed
}

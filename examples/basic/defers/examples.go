package defers

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
	m["defer_basic_usage"] = Basic
	m["defer_in_loop"] = DeferInLoop
	m["defer_in_loop_manyfiles"] = DeferInLoopManyFiles
}

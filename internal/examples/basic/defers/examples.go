package defers

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
	m["defer_basic_usage"] = Basic
	m["defer_in_loop"] = DeferInLoop
	m["defer_in_loop_manyfiles"] = DeferInLoopManyFiles
}

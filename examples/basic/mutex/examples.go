package mutex

import (
	"github.com/devlights/try-golang/examples/basic/mutex/nomutex"
	"github.com/devlights/try-golang/examples/basic/mutex/usechannel"
	"github.com/devlights/try-golang/examples/basic/mutex/usemutex"
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
	m["mutex_nomutex"] = nomutex.NoMutex
	m["mutex_usemutex"] = usemutex.UseMutex
	m["mutex_usechannel"] = usechannel.UseChannel
}

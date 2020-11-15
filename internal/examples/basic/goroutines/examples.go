package goroutines

import (
	"github.com/devlights/try-golang/pkg/mappings"
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
	m["goroutines_nonstop"] = NonStop
	m["goroutines_with_done_channel"] = WithDoneChannel
	m["goroutines_with_waitgroup"] = WithWaitGroup
	m["goroutines_with_context_cancel"] = WithContextCancel
	m["goroutines_with_context_timeout"] = WithContextTimeout
}

package goroutines

import (
	"github.com/devlights/try-golang/examples/basic/goroutines/leak"
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
	m["goroutines_nonstop"] = NonStop
	m["goroutines_with_done_channel"] = WithDoneChannel
	m["goroutines_with_waitgroup"] = WithWaitGroup
	m["goroutines_with_context_cancel"] = WithContextCancel
	m["goroutines_with_context_timeout"] = WithContextTimeout
	m["goroutines_select_nil_chan_1"] = SelectNilChan1
	m["goroutines_select_nil_chan_2"] = SelectNilChan2
	m["goroutines_using_chan_semaphore"] = UsingChanSemaphore
	m["goroutines_using_mutex"] = UsingMutex
	m["goroutines_with_context_deadline"] = WithContextDeadline
	m["goroutines_workerpool"] = WorkerPool

	leak.NewRegister().Regist(m)
}

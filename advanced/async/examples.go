package async

import "github.com/devlights/try-golang/interfaces"

type (
	register struct{}
)

func NewRegister() interfaces.Register {
	r := new(register)
	return r
}

func (r *register) Regist(m interfaces.ExampleMapping) {
	m["async01"] = Async01
	m["async_producer_consumer"] = ProducerConsumer
}

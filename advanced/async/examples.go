package async

import (
	"github.com/devlights/try-golang/mappings"
)

type (
	register struct{}
)

func NewRegister() mappings.Register {
	r := new(register)
	return r
}

func (r *register) Regist(m mappings.ExampleMapping) {
	m["async01"] = Async01
	m["async_producer_consumer"] = ProducerConsumer
}

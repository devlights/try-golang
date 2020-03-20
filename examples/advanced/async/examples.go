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
	m["async_dir_walk_recursive"] = DirWalkRecursive
	m["async_take_first_10items"] = TakeFirst10Items
}

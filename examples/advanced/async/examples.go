package async

import (
	"github.com/devlights/try-golang/mappings"
)

type (
	register struct{}
)

// NewRegister -- このパッケージ用のサンプルを登録する mappings.Register を生成します。
func NewRegister() mappings.Register {
	r := new(register)
	return r
}

// Regist -- サンプルを登録します。
func (r *register) Regist(m mappings.ExampleMapping) {
	m["async01"] = Async01
	m["async_producer_consumer"] = ProducerConsumer
	m["async_dir_walk_recursive"] = DirWalkRecursive
	m["async_take_first_10items"] = TakeFirst10Items
	m["async_ordone_one_input"] = OrDoneOneInput
}

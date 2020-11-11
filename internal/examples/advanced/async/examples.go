package async

import (
	"github.com/devlights/try-golang/internal/examples/advanced/async/concat"
	"github.com/devlights/try-golang/internal/examples/advanced/async/fanin"
	"github.com/devlights/try-golang/internal/examples/advanced/async/ordone"
	"github.com/devlights/try-golang/pkg/mappings"
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
	m["async_ordone_one_input"] = ordone.OneInput
	m["async_ordone_multi_input"] = ordone.MultiInput
	m["async_multi_channel_concat"] = concat.MultiChannelConcat
	m["async_multi_channel_fanin"] = fanin.MultiChannelFanIn
	m["async_ordered_after_async_proc"] = OrderedAfterAsyncProc
}

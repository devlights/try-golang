package loops

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
	m["loops_basic_for_loop"] = BasicForLoop
	m["loops_basic_foreach"] = BasicForeach
	m["loops_channel_loop"] = ChannelLoop
	m["loops_map_loop"] = MapLoop
	m["loops_range_loop"] = RangeLoop
}

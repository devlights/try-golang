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
	m["loops_while_loop"] = WhileLoop
	m["loops_channel_loop"] = ChannelLoop
	m["loops_slice_loop"] = SliceLoop
	m["loops_map_loop"] = MapLoop
	m["loops_range_loop"] = RangeLoop
	m["loops_infinite_loop"] = InfiniteLoop
	m["loops_two_variable"] = ForLoopTwoVariable
	m["loops_for_loop_copy_value"] = ForLoopCopyValue
}

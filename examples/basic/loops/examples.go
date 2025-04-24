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
	m["loops_diff_range_and_normal"] = DiffRangeLoopAndNormalLoop
	m["loops_range_loop_tmpvalue_with_array"] = RangeLoopTmpValueWithArray
	m["loops_go122_loop_variable"] = Go122LoopVariable
	m["loops_go122_range_over_integer"] = Go122RangeOverInterger
	m["loops_dowhile"] = DoWhile
}

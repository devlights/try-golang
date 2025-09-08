package sliceop

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
	m["sliceop_basic01"] = Basic01
	m["sliceop_basic02"] = Basic02
	m["sliceop_basic03"] = Basic03
	m["sliceop_basic04"] = Basic04
	m["sliceop_basic05"] = Basic05
	m["sliceop_reverse"] = Reverse
	m["sliceop_append"] = Append
	m["sliceop_pointer"] = Pointer
	m["sliceop_copy"] = Copy
	m["sliceop_clear"] = Clear
	m["sliceop_deep_equal"] = DeepEqual
	m["sliceop_concat"] = Concat
	m["sliceop_remove_all_elements"] = RemoveAllElements
	m["sliceop_keep_allocated_memory"] = KeepAllocatedMemory
	m["sliceop_nil_append"] = NilAppend
	m["sliceop_three_index"] = ThreeIndex
	m["sliceop_declare_empty_slice"] = DeclareEmtpySlice
	m["sliceop_convert_to_array_go117"] = ConvertToArrayGo117
	m["sliceop_append_special_behavior"] = AppendSpecialBehavior
	m["sliceop_iter_values"] = IterValues
	m["sliceop_iter_all"] = IterAll
	m["sliceop_iter_chunk"] = Chunk
	m["sliceop_normal_fullslice_copy"] = NormalFullsliceCopy
}

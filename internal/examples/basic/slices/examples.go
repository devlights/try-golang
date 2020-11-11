package slices

import (
	"github.com/devlights/try-golang/mappings"
)

type (
	register struct{}
)

// NewRegister -- このパッケージ用のサンプルを登録する mappings.Register を生成します。
func NewRegister() mappings.Register {
	return new(register)
}

// Regist -- 登録します.
func (r *register) Regist(m mappings.ExampleMapping) {
	m["slice_basic01"] = Basic01
	m["slice_basic02"] = Basic02
	m["slice_basic03"] = Basic03
	m["slice_basic04"] = Basic04
	m["slice_basic05"] = Basic05
	m["slice_reverse"] = Reverse
	m["slice_append"] = Append
	m["slice_pointer"] = Pointer
	m["slice_copy"] = Copy
	m["slice_clear"] = Clear
	m["slice_deep_equal"] = DeepEqual
	m["slice_concat"] = Concat
	m["slice_remove_all_elements"] = RemoveAllElements
	m["slice_keep_allocated_memory"] = KeepAllocatedMemory
	m["slice_nil_append"] = NilAppend
}

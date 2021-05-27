package structs

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
	m["struct_basic01"] = Basic01
	m["struct_basic02"] = Basic02
	m["struct_basic03"] = Basic03
	m["struct_basic04"] = Basic04
	m["struct_anonymous_struct"] = StructAnonymousStruct
	m["struct_empty_struct"] = EmptyStruct
	m["struct_deep_equal"] = StructDeepEqual
	m["struct_blank_identifier"] = BlankIdentifier
	m["struct_same_method"] = SameMethodOnEachTypes
	m["struct_memory_padding"] = MemoryPadding
}

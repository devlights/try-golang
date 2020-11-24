package structs

import (
	"github.com/devlights/try-golang/pkg/mappings"
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
	m["struct_basic01"] = Basic01
	m["struct_basic02"] = Basic02
	m["struct_basic03"] = Basic03
	m["struct_basic04"] = Basic04
	m["struct_anonymous_struct"] = StructAnonymousStruct
	m["struct_empty_struct"] = EmptyStruct
	m["struct_deep_equal"] = StructDeepEqual
	m["struct_blank_identifier"] = BlankIdentifier
	m["struct_same_method"] = SameMethodOnEachTypes
}
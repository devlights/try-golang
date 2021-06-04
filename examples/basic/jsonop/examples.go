package jsonop

import "github.com/devlights/try-golang/mapping"

type (
	register struct{}
)

var (
	_ mapping.Register = (*register)(nil)
)

func NewRegister() mapping.Register {
	return new(register)
}

func (*register) Regist(m mapping.ExampleMapping) {
	m["json_marshal_non_indent"] = MarshalNonIndent
	m["json_marshal_indent"] = MarshalIndent
	m["json_marshal_slice"] = MarshalSlice
	m["json_marshal_map"] = MarshalMap
	m["json_unmarshal_struct"] = UnmarshalStruct
	m["json_unmarshal_slice"] = UnmarshalSlice
	m["json_unmarshal_map"] = UnmarshalMap
}

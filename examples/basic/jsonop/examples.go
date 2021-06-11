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
	m["json_marshal_date_rfc3339"] = MarshalDateRfc3339
	m["json_marshal_date_custom"] = MarshalDateCustom
	m["json_unmarshal_struct"] = UnmarshalStruct
	m["json_unmarshal_slice"] = UnmarshalSlice
	m["json_unmarshal_map"] = UnmarshalMap
	m["json_unmarshal_date_rfc3339"] = UnmarshalDateRfc3339
	m["json_unmarshal_date_custom"] = UnmarshalDateCustom
	m["json_decoder"] = Decoder
}

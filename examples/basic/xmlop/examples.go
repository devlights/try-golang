package xmlop

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
	m["xml_marshal"] = Marshal
	m["xml_marshal_indent"] = MarshalIndent
	m["xml_unmarshal"] = Unmarshal
	m["xml_decoder"] = Decoder
}

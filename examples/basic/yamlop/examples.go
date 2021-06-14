package yamlop

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
	m["yaml_marshal"] = Marshal
	m["yaml_unmarshal"] = Unmarshal
	m["yaml_decoder"] = Decoder
	m["yaml_encoder"] = Encoder
}

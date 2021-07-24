package base64op

import "github.com/devlights/try-golang/mapping"

type (
	register struct{}
)

func NewRegister() mapping.Register {
	return new(register)
}

func (r *register) Regist(m mapping.ExampleMapping) {
	m["base64op_encode"] = Encode
	m["base64op_decode"] = Decode
}

package crypto

import "github.com/devlights/try-golang/interfaces"

type (
	register struct{}
)

func NewRegister() interfaces.Register {
	r := new(register)
	return r
}

func (r *register) Regist(m interfaces.ExampleMapping) {
	m["crypto_bcrypt_password_hash"] = BcryptPasswordHash
}

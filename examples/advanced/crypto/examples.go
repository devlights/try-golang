package crypto

import (
	"github.com/devlights/try-golang/mappings"
)

type (
	register struct{}
)

func NewRegister() mappings.Register {
	r := new(register)
	return r
}

func (r *register) Regist(m mappings.ExampleMapping) {
	m["crypto_bcrypt_password_hash"] = BcryptPasswordHash
}

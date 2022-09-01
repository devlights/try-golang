package cryptos

import (
	"github.com/devlights/try-golang/examples/basic/cryptos/aes"
	"github.com/devlights/try-golang/examples/basic/cryptos/bcrypt"
	"github.com/devlights/try-golang/examples/basic/cryptos/checksum"
	"github.com/devlights/try-golang/examples/basic/cryptos/rand"
	"github.com/devlights/try-golang/mapping"
)

type (
	register struct{}
)

// NewRegister -- このパッケージ用のサンプルを登録する mapping.Register を生成します。
func NewRegister() mapping.Register {
	return new(register)
}

// Regist -- 登録します.
func (r *register) Regist(m mapping.ExampleMapping) {
	m["crypto_md5_checksum"] = checksum.Md5Checksum
	m["crypto_aes_ecb"] = aes.Ecb
	m["crypto_aes_cbc"] = aes.Cbc
	m["crypto_rand_reader"] = rand.Reader
	m["crypto_rand_read"] = rand.Read
	m["crypto_bcrypt_generate"] = bcrypt.Generate
	m["crypto_bcrypt_compare"] = bcrypt.Compare
}

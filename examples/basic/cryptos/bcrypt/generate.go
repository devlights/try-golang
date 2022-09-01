package bcrypt

import (
	"strings"

	"github.com/devlights/gomy/output"
	"golang.org/x/crypto/bcrypt"
)

// Generate -- golang.org/x/crypto/bcrypt を使って bcrypt パスワードハッシュ を生成するサンプルです.
//
// # REFERNCES
//   - https://pkg.go.dev/golang.org/x/crypto/bcrypt
//   - https://zenn.dev/kou_pg_0131/articles/go-digest-and-compare-by-bcrypt
//   - https://medium-company.com/bcrypt/
func Generate() error {
	var (
		b      = []byte("helloworld")
		hashed []byte
		err    error
	)

	hashed, err = bcrypt.GenerateFromPassword(b, bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	output.Stdoutf("[bcrypt][original]", "%s\n", b)
	output.Stdoutf("[bcrypt][hashed  ]", "%s\n", hashed)

	var (
		s           = string(hashed)
		parts       = strings.Split(s, "$")
		hashAlgo    = parts[1]
		streatching = parts[2]
		salt        = parts[3][:22]
		hash        = parts[3][22:]
	)
	output.Stdoutl("[bcrypt][parts][algorithm  ]", hashAlgo)
	output.Stdoutl("[bcrypt][parts][streatching]", streatching)
	output.Stdoutl("[bcrypt][parts][salt       ]", salt)
	output.Stdoutl("[bcrypt][parts][hash       ]", hash)

	return nil
}

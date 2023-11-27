package bcrypt

import (
	"errors"

	"github.com/devlights/gomy/output"
	"golang.org/x/crypto/bcrypt"
)

// Compare --  golang.org/x/crypto/bcrypt を使って生成したパスワードハッシュと比較するサンプルです.
//
// # REFERNCES
//   - https://pkg.go.dev/golang.org/x/crypto/bcrypt
//   - https://zenn.dev/kou_pg_0131/articles/go-digest-and-compare-by-bcrypt
//   - https://medium-company.com/bcrypt/
func Compare() error {
	var (
		b      = []byte("helloworld")
		hashed []byte
		err    error
	)

	hashed, err = bcrypt.GenerateFromPassword(b, bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	var (
		b2 = []byte("helloworld")
		b3 = []byte("helloworlD")
	)

	for _, v := range [][]byte{b2, b3} {

		err = bcrypt.CompareHashAndPassword(hashed, v)
		if err != nil {
			if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
				output.Stdoutf("[mismatch]", "%s\n", v)
				continue
			}

			return err
		}

		output.Stdoutf("[match]", "%s\n", v)
	}

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: crypto_bcrypt_compare

	   [Name] "crypto_bcrypt_compare"
	   [match]              helloworld
	   [mismatch]           helloworlD


	   [Elapsed] 180.46252ms
	*/

}

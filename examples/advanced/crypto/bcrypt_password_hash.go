package crypto

import (
	"syscall"

	"github.com/devlights/gomy/output"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/term"
)

// BcryptPasswordHash は、 bcrypt パッケージを利用してパスワードのハッシュ化をしてみるサンプルです.
// パスワードの読み取りには、 terminal パッケージの ReadPassword() を利用しています.
//
// REFERENCES::
//   - https://medium.com/@jcox250/password-hash-salt-using-golang-b041dc94cb72
//   - https://pkg.go.dev/golang.org/x/crypto/bcrypt?tab=doc
//   - https://pkg.go.dev/golang.org/x/crypto/ssh/terminal?tab=doc
//   - https://ja.wikipedia.org/wiki/Bcrypt
//   - https://liginc.co.jp/377191
//   - https://stackoverflow.com/questions/30363790/silence-user-input-in-scan-function
func BcryptPasswordHash() error {

	// パスワード読み取り
	password, err := readPassword()
	if err != nil {
		return err
	}

	// bcrypt で ハッシュ化
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// ハッシュ化したものを文字列で表示
	output.Stdoutl("Hashed Password: ", string(hashedPassword))

	// 再度パスワード読み取り
	password2, err := readPassword()
	if err != nil {
		return err
	}

	// 一致するかを bcrypt パッケージの関数で確認
	err = bcrypt.CompareHashAndPassword(hashedPassword, password2)
	if err != nil {
		output.Stdoutl("CompareHashAndPassword", "一致しませんでした")
	} else {
		output.Stdoutl("CompareHashAndPassword", "一致しました")
	}

	return nil
}

func readPassword() ([]byte, error) {

	output.Stdoutf("ENTER Password: ", "")

	password, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return nil, err
	}
	output.Stdoutl("", "")

	return password, nil
}

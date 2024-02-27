package osop

import (
	"os"

	"github.com/devlights/gomy/output"
)

// ExpandEnv は、os.ExpandEnv() のサンプルです。
//
// ExpandEnv は、現在の環境変数の値に従って、文字列中の ${var} または $var を置き換えます。
// 未定義の変数への参照は空文字列に置き換えられます。
//
// 戻り値は string で、error は返らない。
//
// # REFERENCES
//
//   - https://pkg.go.dev/os@go1.22.0#ExpandEnv
func ExpandEnv() error {
	var (
		env1 = os.ExpandEnv("${HOME}")       // ${VAL}形式
		env2 = os.ExpandEnv("$SONZAISHINAI") // $VAL形式
		env3 = os.ExpandEnv("home is ${HOME}, hostname is ${HOSTNAME}")
	)

	output.Stdoutf("[env1]", "%q\n", env1)
	output.Stdoutf("[env2]", "%q\n", env2)
	output.Stdoutf("[env3]", "%q\n", env3)

	return nil
}

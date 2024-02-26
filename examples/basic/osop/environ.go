package osop

import (
	"os"
	"strings"

	"github.com/devlights/gomy/output"
)

// Environ は、os.Environ()のサンプルです。
//
// os.Environ() は、現在の環境変数の値を key=value 形式の文字列で返す。
// 戻り値は []string 。
//
// # REFERENCES
//   - https://pkg.go.dev/os@go1.22.0#Environ
func Environ() error {
	for _, env := range os.Environ() {
		var (
			kv = strings.Split(env, "=")
			k  = kv[0]
			v  = kv[1]
		)

		if !strings.HasPrefix(k, "H") {
			continue
		}

		output.Stdoutl("[env]", k, v)
	}

	return nil
}

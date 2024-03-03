package osop

import (
	"os"

	"github.com/devlights/gomy/output"
)

// Clearenv は、os.Clearenv() のサンプルです。
//
// 全環境変数をクリアします。(このプロセス上での)
//
// # REFERENCES
//
//   - https://pkg.go.dev/os@go1.22.0#Clearenv
func Clearenv() error {
	var envs []string

	envs = os.Environ()
	output.Stdoutl("[os.Environ]", len(envs))

	os.Clearenv()

	envs = os.Environ()
	output.Stdoutl("[os.Environ]", len(envs))

	return nil
}

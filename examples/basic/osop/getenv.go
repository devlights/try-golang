package osop

import (
	"os"

	"github.com/devlights/gomy/output"
)

// GetEnv は、os.GetEnv() のサンプルです。
//
// Getenvは、キーで指定された環境変数の値を取得します。
// 値が返されるが、変数が存在しない場合は空が返ります。
// 空の値と未設定の値を区別するには、LookupEnvを使用します。
//
// 戻り値は string で、error は返却されない。
//
// # REFERENCES
//
//   - https://pkg.go.dev/os@go1.22.0#Getenv
func GetEnv() error {
	const (
		ENV1 = "HOSTNAME"
		ENV2 = "SONZAISHINAIKEY"
	)

	var (
		env1 = os.Getenv(ENV1)
		env2 = os.Getenv(ENV2)
	)

	//
	// env2 の方は存在しない環境変数のため空が返る。
	// この「空」の値が、「存在しない環境変数」なのか「存在するが値が空」なのかを
	// 見極める必要がある場合は、os.LookupEnv() の方を使う。
	//
	output.Stdoutl("[env1]", env1)
	output.Stdoutl("[env2]", env2)

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: osop_getenv

	   [Name] "osop_getenv"
	   [env1]               devlights-trygolang-q7kq6quld1n
	   [env2]


	   [Elapsed] 43.03µs
	*/

}

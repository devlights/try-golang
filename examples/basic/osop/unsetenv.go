package osop

import (
	"os"

	"github.com/devlights/gomy/output"
)

// Unsetenv は、os.Unsetenv() のサンプルです。
//
// 指定された環境変数の値をクリアします。
// 一時的な環境変数を用意する際に、os.Setenv()とペアで以下のように
// よく利用される。
//
//	os.Setenv("MYENV", "HELLOWORLD")
//	defer os.Unsetenv("MYENV")
//
// # REFERENCES
//
//   - https://pkg.go.dev/os@go1.22.0#Unsetenv
func Unsetenv() error {
	const (
		ENVKEY = "MYENV"
		ENVVAL = "HELLOWORLD"
	)

	var (
		env string
		ok  bool
		err error
	)

	err = os.Setenv(ENVKEY, ENVVAL)
	if err != nil {
		return err
	}

	env, ok = os.LookupEnv(ENVKEY)
	output.Stdoutf("[MYENV(before)]", "VALUE=%q\tOK=%v\n", env, ok)

	err = os.Unsetenv(ENVKEY)
	if err != nil {
		return err
	}

	env, ok = os.LookupEnv(ENVKEY)
	output.Stdoutf("[MYENV(after )]", "VALUE=%q\tOK=%v\n", env, ok)

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: osop_unsetenv

	   [Name] "osop_unsetenv"
	   [MYENV(before)]      VALUE="HELLOWORLD" OK=true
	   [MYENV(after )]      VALUE=""   OK=false


	   [Elapsed] 70.13µs
	*/

}

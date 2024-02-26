package osop

import (
	"os"

	"github.com/devlights/gomy/output"
)

// LookupEnv は、os.LookupEnv() のサンプルです。
//
// LookupEnv は、キーで指定された環境変数の値を取得します。
// その変数が環境に存在する場合、その値 (空でもよい) が返され、ブール値は true になります。
// そうでない場合は、返される値は空で、ブール値は false になります。
//
// # REFERENCES
//
//   - https://pkg.go.dev/os@go1.22.0#LookupEnv
func LookupEnv() error {
	const (
		ENV1 = "HOSTNAME"
		ENV2 = "SONZAISHINAIKEY"
	)

	var (
		v  string
		ok bool
		p  = func(prefix string, v string, ok bool) {
			if ok {
				output.Stdoutl(prefix, v)
			} else {
				output.Stdoutl(prefix, "not found")
			}
		}
	)

	v, ok = os.LookupEnv(ENV1)
	p("[ENV1]", v, ok)

	v, ok = os.LookupEnv(ENV2)
	p("[ENV2]", v, ok)

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: osop_lookupenv

	   [Name] "osop_lookupenv"
	   [ENV1]               devlights-trygolang-q7kq6quld1n
	   [ENV2]               not found


	   [Elapsed] 19.3µs
	*/

}

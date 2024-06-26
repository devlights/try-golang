package flags

import (
	"flag"
	"strings"

	"github.com/devlights/gomy/output"
)

type (
	// flag.Var() で指定するホスト名リスト。カンマ区切り。
	hosts []string
	// flag.Var() で指定するパスリスト。コロン区切り。
	paths []string
)

func (h hosts) String() string {
	output.Stderrl("[CALL]", "flag.Value.String()")

	return strings.Join(h, ",")
}

func (h *hosts) Set(v string) error {
	output.Stderrl("[CALL]", "flag.Value.Set()")

	*h = strings.Split(v, ",")
	return nil
}

func (p paths) String() string {
	return strings.Join(p, ":")
}

func (p *paths) Set(v string) error {
	*p = strings.Split(v, ":")
	return nil
}

// Var は、flag.Var() のサンプルです。
//
// flag.Var() には、任意の値をフラグハンドリングのための値として指定出来ます。
// 条件として、flag.Valueインターフェースを実装している必要があります。
//
// # REFERENCES
//   - https://pkg.go.dev/flag@go1.22.4#Var
//   - https://pkg.go.dev/flag@go1.22.4#Value
//   - https://serverfault.com/a/846523
//   - https://www.rfc-editor.org/rfc/rfc6761
func Var() error {
	var (
		hs    hosts
		ps    paths
		fs    = flag.NewFlagSet("", flag.ExitOnError)
		logfn = func(message string, fn func()) {
			output.Stderrl("[CALL]", message)
			fn()
		}
	)

	logfn("fs.Var", func() { fs.Var(&hs, "hosts", "host names. comma separated.") })
	fs.Var(&ps, "paths", "path list. coron separated.")

	logfn("fs.Parse", func() {
		fs.Parse([]string{"-hosts", "example.invalid,dummy.local,my.test", "-paths", "/path/to/a:/bin:/usr/bin"})
	})

	for _, h := range hs {
		output.Stdoutl("[h]", h)
	}

	for _, p := range ps {
		output.Stdoutl("[p]", p)
	}

	return nil

	/*
		$ task
		task: [build] go build .
		task: [run] ./try-golang -onetime

		ENTER EXAMPLE NAME: flags_var

		[Name] "flags_var"
		[CALL]               fs.Var
		[CALL]               flag.Value.String()
		[CALL]               fs.Parse
		[CALL]               flag.Value.Set()
		[h]                  example.invalid
		[h]                  dummy.local
		[h]                  my.test
		[p]                  /path/to/a
		[p]                  /bin
		[p]                  /usr/bin


		[Elapsed] 86µs
	*/

}

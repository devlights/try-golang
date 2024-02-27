package osop

import (
	"os"

	"github.com/devlights/gomy/output"
)

// Expand は、os.Expand() のサンプルです。
//
// Expandは、マッピング関数に基づいて文字列の${var}または$varを置き換えます。
// マッピング関数の書式は
//
//	func(string) string
//
// となっています。
// os.ExpandEnv() は、以下と同じことになります。
//
//	os.Expand(s, os.Getenv)
//
// # REFERENCES
//
//   - https://pkg.go.dev/os@go1.22.0#Expand
func Expand() error {
	var (
		fn = func(s string) string {
			return "helloworld"
		}
		v = os.Expand("${HI}", fn)
	)

	output.Stdoutl("[HI]", v)
	output.Stdoutl("[HOME]", os.Expand("${HOME}", os.Getenv))

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: osop_expand

	   [Name] "osop_expand"
	   [HI]                 helloworld
	   [HOME]               /home/gitpod


	   [Elapsed] 11.04µs
	*/

}

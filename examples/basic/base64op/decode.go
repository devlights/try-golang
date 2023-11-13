package base64op

import (
	"bytes"
	"encoding/base64"

	"github.com/devlights/gomy/output"
)

// Decode -- base64.StdEncoding.Decode のサンプルです.
//
// REFERENCES:
//   - https://pkg.go.dev/encoding/base64
//   - https://pkg.go.dev/net/http
//   - https://gobyexample.com/base64-encoding
//   - https://golangdocs.com/base64-encoding-decoding-golang
//   - https://golangdocs.com/golang-download-files
//   - https://qiita.com/daijuk/items/d5c52b780e90387f2390
func Decode() error {
	const (
		//         helloworld
		encoded = "aGVsbG93b3JsZA=="
	)

	var (
		src = bytes.NewBufferString(encoded)
		dst = make([]byte, base64.StdEncoding.DecodedLen(src.Len()))
		err error
	)

	if _, err = base64.StdEncoding.Decode(dst, src.Bytes()); err != nil {
		return err
	}

	output.Stdoutl("[original]", encoded)
	output.Stdoutl("[decode  ]", string(dst))

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: base64op_decode

	   [Name] "base64op_decode"
	   [original]           aGVsbG93b3JsZA==
	   [decode  ]           helloworld


	   [Elapsed] 34.52µs
	*/

}

package base64op

import (
	"bytes"
	"encoding/base64"

	"github.com/devlights/gomy/output"
)

// Encode -- base64.StdEncoding.Encode のサンプルです.
//
// REFERENCES:
//   - https://pkg.go.dev/encoding/base64
//   - https://pkg.go.dev/net/http
//   - https://gobyexample.com/base64-encoding
//   - https://golangdocs.com/base64-encoding-decoding-golang
//   - https://golangdocs.com/golang-download-files
//   - https://qiita.com/daijuk/items/d5c52b780e90387f2390
func Encode() error {
	const (
		imgUrl = "https://unsplash.com/photos/Cv96uQrD5HI/download?force=true"
	)

	var (
		dl  = downloder(imgUrl)
		buf *bytes.Buffer
		err error
	)

	if buf, err = dl.download(); err != nil {
		return err
	}
	output.Stdoutl("[original length      ]", buf.Len())

	var (
		src = buf.Bytes()
		dst = make([]byte, base64.StdEncoding.EncodedLen(len(src)))
	)

	base64.StdEncoding.Encode(dst, src)

	output.Stdoutl("[base64 encoded length]", len(dst))
	output.Stdoutl("[incremental rate     ]", float32(len(dst))/float32(len(src)))
	output.Stdoutl("[first 100 bytes      ]", string(dst[:99]))

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: base64op_encode

	   [Name] "base64op_encode"
	   [original length      ] 1062765
	   [base64 encoded length] 1417020
	   [incremental rate     ] 1.3333334
	   [first 100 bytes      ] /9j/4AAQSkZJRgABAQEASABIAAD/4gxYSUNDX1BST0ZJTEUAAQEAAAxITGlubwIQAABtbnRyUkdCIFhZWiAHzgACAAkABgAxAAB


	   [Elapsed] 1.985106023s
	*/

}

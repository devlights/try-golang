package ioop

import (
	"bytes"
	"compress/gzip"
	"encoding/hex"
	"hash/crc32"
	"io"
	"os"

	"github.com/devlights/gomy/output"
)

// MultiWrite は、io.MultiWriterを利用してgzip圧縮しながらCRCチェックサムも算出するサンプルです.
//
// # REFERENCES
//
//   - https://pkg.go.dev/io@go1.22.2#MultiWriter
//   - https://pkg.go.dev/hash/crc32@go1.22.2#NewIEEE
//   - https://pkg.go.dev/compress/gzip@go1.22.2#Writer
//   - https://pkg.go.dev/encoding/hex@go1.22.2#Dumper
func MultiWrite() error {
	var (
		data   = []byte("hello world こんにちは 世界")
		buf    = new(bytes.Buffer)
		gzipW  = gzip.NewWriter(buf)
		crcW   = crc32.NewIEEE()
		hexW   = hex.Dumper(os.Stdout)
		writer = io.MultiWriter(gzipW, crcW, hexW)
		err    error
	)
	defer gzipW.Close()

	_, err = writer.Write(data)
	if err != nil {
		return err
	}
	hexW.Close()

	output.Stdoutf("[orig]", "%x\n", data)
	output.Stdoutf("[gzip]", "%x\n", buf.Bytes())
	output.Stdoutf("[crc ]", "%x\n", crcW.Sum32())

	return nil

	/*
		$ task
		task: [build] go build .
		task: [run] ./try-golang -onetime

		ENTER EXAMPLE NAME: ioop_multiwrite

		[Name] "ioop_multiwrite"
		00000000  68 65 6c 6c 6f 20 77 6f  72 6c 64 20 e3 81 93 e3  |hello world ....|
		00000010  82 93 e3 81 ab e3 81 a1  e3 81 af 20 e4 b8 96 e7  |........... ....|
		00000020  95 8c                                             |..|
		[orig]               68656c6c6f20776f726c6420e38193e38293e381abe381a1e381af20e4b896e7958c
		[gzip]               1f8b08000000000000ff
		[crc ]               6535a281


		[Elapsed] 394.39µs
	*/

}

package byteop

import (
	"bytes"
	"errors"

	"github.com/devlights/gomy/output"
)

// CutPrefixSuffix は、Go 1.20 で追加された bytes.{CutPrefix,CutSuffix} のサンプルです.
//
// bytes.{CutPrefix,CutSuffix} と strings.{CutPrefix,CutSuffix} は対で追加されている。
//
// # REFERENCES
//   - https://pkg.go.dev/bytes@go1.20.2#CutPrefix
//   - https://pkg.go.dev/bytes@go1.20.2#CutSuffix
func CutPrefixSuffix() error {
	var (
		prefix = []byte("hello")
		suffix = []byte("world")
		body   = []byte("12345")
		sep    = []byte{}
		data   = bytes.Join([][]byte{prefix, body, suffix}, sep)
	)

	cut1, found := bytes.CutPrefix(data, prefix)
	if !found {
		return errors.New("prefix is not found")
	}

	cut2, found := bytes.CutSuffix(data, suffix)
	if !found {
		return errors.New("suffix is not found")
	}

	output.Stdoutl("[CutPrefix]", cut1, string(cut1))
	output.Stdoutl("[CutSuffix]", cut2, string(cut2))

	return nil

	/*
	   $ task
	   task: Task "build" is up to date
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: byteop_cut_prefix_suffix

	   [Name] "byteop_cut_prefix_suffix"
	   [CutPrefix]          [49 50 51 52 53 119 111 114 108 100] 12345world
	   [CutSuffix]          [104 101 108 108 111 49 50 51 52 53] hello12345


	   [Elapsed] 53.02µs
	*/

}

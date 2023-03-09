package strs

import (
	"errors"
	"strings"

	"github.com/devlights/gomy/output"
)

// CutPrefixSuffix は、Go 1.20 で追加された strings.{CutPrefix,CutSuffix} のサンプルです.
//
// bytes.{CutPrefix,CutSuffix} と strings.{CutPrefix,CutSuffix} は対で追加されている。
//
// # REFERENCES
//   - https://pkg.go.dev/strings@go1.20.2#CutPrefix
//   - https://pkg.go.dev/strings@go1.20.2#CutSuffix
func CutPrefixSuffix() error {
	var (
		prefix = "hello"
		suffix = "world"
		body   = "12345"
		sep    = ""
		data   = strings.Join([]string{prefix, body, suffix}, sep)
	)

	cut1, found := strings.CutPrefix(data, prefix)
	if !found {
		return errors.New("prefix is not found")
	}

	cut2, found := strings.CutSuffix(data, suffix)
	if !found {
		return errors.New("suffix is not found")
	}

	output.Stdoutl("[CutPrefix]", cut1)
	output.Stdoutl("[CutSuffix]", cut2)

	return nil
}

package jsonop

import (
	"encoding/json"

	"github.com/devlights/gomy/output"
)

// MarshalIndent は、json.MarshalIndent() を利用したサンプルです.
func MarshalIndent() error {
	type (
		V struct {
			Value1 string `json:"name"`
			Value2 string `json:"name2"`
		}
	)

	var (
		v = V{
			Value1: "hello",
			Value2: "world",
		}
	)

	var (
		buf []byte
		err error
	)

	if buf, err = json.MarshalIndent(&v, "", "  "); err != nil {
		return err
	}

	output.Stdoutl("[marshal]", string(buf))

	return nil
}

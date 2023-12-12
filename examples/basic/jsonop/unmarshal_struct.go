package jsonop

import (
	"bytes"
	"encoding/json"

	"github.com/devlights/gomy/output"
)

// UnmarshalStruct は、json.Unmarshal() を利用したサンプルです.
func UnmarshalStruct() error {
	type (
		V struct {
			Value1 string `json:"name"`
			Value2 string `json:"value"`
		}
	)

	var (
		buf = bytes.NewBufferString(`
		{
			"name": "value1",
			"value": "value2"
		}
		`)
	)

	var (
		v   V
		err error
	)

	if err = json.Unmarshal(buf.Bytes(), &v); err != nil {
		return err
	}

	output.Stdoutl("[original]", buf.String())
	output.Stdoutf("[unmarshal]", "%#v\n", v)

	return nil
}

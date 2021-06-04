package jsonop

import (
	"bytes"
	"encoding/json"

	"github.com/devlights/gomy/output"
	"github.com/devlights/try-golang/examples/basic/jsonop/types"
)

// UnmarshalDateCustom は、独自の日付文字列 を json.Unmarshal した場合のサンプルです.
func UnmarshalDateCustom() error {
	var (
		buf = bytes.NewBufferString(`"2021/06/04"`)
	)

	var (
		v   types.YyyyMmDd
		err error
	)

	if err = json.Unmarshal(buf.Bytes(), &v); err != nil {
		return err
	}

	output.Stdoutl("[original]", buf.String())
	output.Stdoutl("[unmarshal]", v)

	return nil
}

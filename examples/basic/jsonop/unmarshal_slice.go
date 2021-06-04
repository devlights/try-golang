package jsonop

import (
	"bytes"
	"encoding/json"

	"github.com/devlights/gomy/output"
)

// UnmarshalSlice は、json 配列 を json.Unmarshal した場合のサンプルです.
func UnmarshalSlice() error {
	var (
		buf = bytes.NewBufferString(`
			[
				"golang",
				"java",
				"dotnet",
				"python"
			]
		`)
	)

	var (
		items []string
		err   error
	)

	if err = json.Unmarshal(buf.Bytes(), &items); err != nil {
		return err
	}

	output.Stdoutl("[original]", buf.String())
	output.Stdoutl("[unmarshal]", items)

	return nil
}

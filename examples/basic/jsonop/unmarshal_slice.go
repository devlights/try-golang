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

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: json_unmarshal_slice

	   [Name] "json_unmarshal_slice"
	   [original]
	                           [
	                                   "golang",
	                                   "java",
	                                   "dotnet",
	                                   "python"
	                           ]

	   [unmarshal]          [golang java dotnet python]


	   [Elapsed] 57.75µs
	*/

}

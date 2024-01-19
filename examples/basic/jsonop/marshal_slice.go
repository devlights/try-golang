package jsonop

import (
	"encoding/json"

	"github.com/devlights/gomy/output"
)

// MarshalSlice は、json.Marshal() で スライス をマーシャルした場合のサンプルです.
func MarshalSlice() error {
	var (
		items = []string{
			"golang",
			"java",
			"dotnet",
			"python",
		}
	)

	var (
		buf []byte
		err error
	)

	if buf, err = json.Marshal(items); err != nil {
		return err
	}

	output.Stdoutl("[marshal]", string(buf))

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: json_marshal_slice

	   [Name] "json_marshal_slice"
	   [marshal]            ["golang","java","dotnet","python"]


	   [Elapsed] 46.83µs
	*/

}

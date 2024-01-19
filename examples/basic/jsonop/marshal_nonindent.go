package jsonop

import (
	"encoding/json"

	"github.com/devlights/gomy/output"
)

// MarshalNonIndent は、json.Marshal() を利用したサンプルです.
func MarshalNonIndent() error {
	type (
		V struct {
			notSerialize string
			Value1       string `json:"value"`
			Value2       string
		}
	)

	var (
		v = V{
			notSerialize: "not_serialize",
			Value1:       "value1",
			Value2:       "value2",
		}
	)

	var (
		buf []byte
		err error
	)

	if buf, err = json.Marshal(&v); err != nil {
		return err
	}

	output.Stdoutl("[marshal]", string(buf))

	return nil

	/*
	   $ task
	   task: Task "build" is up to date
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: json_marshal_non_indent

	   [Name] "json_marshal_non_indent"
	   [marshal]            {"value":"value1","Value2":"value2"}


	   [Elapsed] 53.83µs
	*/

}

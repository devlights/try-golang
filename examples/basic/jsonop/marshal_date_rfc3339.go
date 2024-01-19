package jsonop

import (
	"encoding/json"
	"time"

	"github.com/devlights/gomy/output"
)

// MarshalDateRfc3339 は、RFC3339形式の日付を json.Marshal するサンプルです.
func MarshalDateRfc3339() error {
	type V struct {
		T time.Time `json:"t"`
	}

	var (
		v = V{time.Now()}
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
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: json_marshal_date_rfc3339

	   [Name] "json_marshal_date_rfc3339"
	   [marshal]            {"t":"2024-01-19T02:06:43.870334862Z"}


	   [Elapsed] 157.78µs
	*/

}

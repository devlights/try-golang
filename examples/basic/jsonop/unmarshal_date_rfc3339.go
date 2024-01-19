package jsonop

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"github.com/devlights/gomy/output"
)

// UnmarshalDateRfc3339 は、RFC3339形式の日付文字列 を json.Unmarshal した場合のサンプルです.
func UnmarshalDateRfc3339() error {
	type V struct {
		T time.Time `json:"t"`
	}

	var (
		t   = time.Now().Format(time.RFC3339)
		buf = bytes.NewBufferString(fmt.Sprintf(`{"t": "%s"}`, t))
	)

	var (
		v   V
		err error
	)

	if err = json.Unmarshal(buf.Bytes(), &v); err != nil {
		return err
	}

	output.Stdoutl("[original]", buf.String())
	output.Stdoutl("[unmarshal]", v)

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: json_unmarshal_date_rfc3339

	   [Name] "json_unmarshal_date_rfc3339"
	   [original]           {"t": "2024-01-19T02:08:43Z"}
	   [unmarshal]          {2024-01-19 02:08:43 +0000 UTC}


	   [Elapsed] 141.589µs
	*/

}

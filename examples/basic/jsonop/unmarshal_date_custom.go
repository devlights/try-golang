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

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: json_unmarshal_date_custom

	   [Name] "json_unmarshal_date_custom"
	   [original]           "2021/06/04"
	   [unmarshal]          2021/06/04


	   [Elapsed] 42.16µs
	*/

}

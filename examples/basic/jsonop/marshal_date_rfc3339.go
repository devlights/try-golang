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
}

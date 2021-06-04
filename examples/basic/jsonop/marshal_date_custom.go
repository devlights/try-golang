package jsonop

import (
	"encoding/json"
	"time"

	"github.com/devlights/gomy/output"
	"github.com/devlights/try-golang/examples/basic/jsonop/types"
)

// MarshalDateCustom は、json.Marshal にて 独自の日付形式 を利用するサンプルです.
func MarshalDateCustom() error {
	var (
		ct = types.YyyyMmDd{Time: time.Now()}
	)

	var (
		buf []byte
		err error
	)

	if buf, err = json.Marshal(&ct); err != nil {
		return err
	}

	output.Stdoutl("[original]", ct)
	output.Stdoutl("[marshal]", string(buf))

	return nil
}

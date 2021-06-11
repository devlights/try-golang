package jsonop

import (
	"bytes"
	"encoding/json"

	"github.com/devlights/gomy/output"
)

// Encoder は、json.Encoder を使ったサンプルです.
func Encoder() error {
	type (
		Message struct {
			Id    int    `json:"id"`
			Value string `json:"value"`
		}
	)

	var (
		buf     = new(bytes.Buffer)
		encoder = json.NewEncoder(buf)
	)

	var (
		msgs = []*Message{
			{Id: 100, Value: "golang"},
			{Id: 200, Value: "flutter"},
		}
	)

	for _, msg := range msgs {
		var (
			err error
		)

		if err = encoder.Encode(msg); err != nil {
			return err
		}
	}

	output.Stdoutf("[encode]", "\n%v\n", buf.String())

	return nil
}

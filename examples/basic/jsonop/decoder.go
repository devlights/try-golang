package jsonop

import (
	"bytes"
	"encoding/json"
	"io"

	"github.com/devlights/gomy/output"
)

// Decoder は、json.NewDecoder を使ったサンプルです.
func Decoder() error {
	const (
		jsonStr = `
		{ "id":100, "value": "golang"  }
		{ "id":200, "value": "flutter" }
		`
	)

	type (
		Message struct {
			Id    int    `json:"id"`
			Value string `json:"value"`
		}
	)

	var (
		reader  = bytes.NewBufferString(jsonStr)
		decoder = json.NewDecoder(reader)
	)

LOOP:
	for {
		var (
			msg Message
			err error
		)

		if err = decoder.Decode(&msg); err != nil {
			switch err {
			case io.EOF:
				break LOOP
			default:
				return err
			}
		}

		output.Stdoutf("[msg]", "%v\n", msg)
	}

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: json_decoder

	   [Name] "json_decoder"
	   [msg]                {100 golang}
	   [msg]                {200 flutter}


	   [Elapsed] 180.72µs
	*/

}

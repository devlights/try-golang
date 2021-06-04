package jsonop

import (
	"bytes"
	"encoding/json"

	"github.com/devlights/gomy/output"
)

// UnmarshalMap は、json マップ を json.Unmarshal した場合のサンプルです.
func UnmarshalMap() error {
	var (
		buf = bytes.NewBufferString(`
		{
			"golang": "fmt.Println",
			"java": "System.out.println",
			"dotnet": "Console.WriteLine",
			"python": "print"
		}
		`)
	)

	var (
		m   map[string]string
		err error
	)

	if err = json.Unmarshal(buf.Bytes(), &m); err != nil {
		return err
	}

	output.Stdoutl("[original]", buf.String())
	output.Stdoutl("[unmarshal]", m)

	return nil
}

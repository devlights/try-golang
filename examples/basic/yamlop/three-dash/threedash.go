package threedash

import (
	"bytes"
	"io"
	"os"

	"github.com/devlights/gomy/output"
	"gopkg.in/yaml.v3"
)

// ThreeDash は、YAML の "---" で複数のドキュメントが存在する場合のサンプルです.
//
// # REFERENCES
//   - https://yaml.org/spec/1.2.2/
func ThreeDash() error {
	// Read file
	var (
		file *os.File
		buf  bytes.Buffer
		err  error
	)
	file, err = os.Open("./examples/basic/yamlop/three-dash/data.yaml")
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(&buf, file)
	if err != nil {
		return err
	}

	// Decode yaml
	var (
		data    []string
		decoder = yaml.NewDecoder(&buf)
	)
	for {
		err = decoder.Decode(&data)
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		output.Stdoutl("[data]", data)
	}

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: yaml_threedash

	   [Name] "yaml_threedash"
	   [data]               [hello world]
	   [data]               [HELLO WORLD]
	   [data]               [WORLD HELLO]


	   [Elapsed] 197.96µs
	*/

}

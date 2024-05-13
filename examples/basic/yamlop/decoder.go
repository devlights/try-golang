package yamlop

import (
	"bytes"

	"github.com/devlights/gomy/output"
	"github.com/devlights/try-golang/examples/basic/yamlop/types"
	"gopkg.in/yaml.v3"
)

// Decoder は、yaml.Decoder のサンプルです.
func Decoder() error {
	const (
		yamlStr = `
languages:
- name: golang
  printfn: fmt.Println
  version:
    major: 1
    minor: 16
- name: java
  printfn: System.out.println
  version:
    major: 16
    minor: 0`
	)

	var (
		buf = bytes.NewBufferString(yamlStr)
		dec = yaml.NewDecoder(buf)
	)

	var (
		v   types.YamlData
		err error
	)

	if err = dec.Decode(&v); err != nil {
		return err
	}

	output.Stdoutf("[decoder]", "%v\n", v)

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: yaml_decoder

	   [Name] "yaml_decoder"
	   [decoder]            {[{golang fmt.Println {1 16}} {java System.out.println {16 0}}]}


	   [Elapsed] 177.78µs
	*/

}

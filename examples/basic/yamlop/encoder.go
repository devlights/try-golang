package yamlop

import (
	"bytes"

	"github.com/devlights/gomy/output"
	"github.com/devlights/try-golang/examples/basic/yamlop/types"
	"gopkg.in/yaml.v3"
)

// Encoder は、yaml.Encoder についてのサンプルです.
func Encoder() error {
	var (
		v = types.YamlData{
			Languages: []types.Language{
				{Name: "golang", PrintFn: "fmt.Println", Version: types.Version{Major: 1, Minor: 16}},
				{Name: "java", PrintFn: "System.out.println", Version: types.Version{Major: 16, Minor: 0}},
			},
		}
	)

	var (
		buf = new(bytes.Buffer)
		enc = yaml.NewEncoder(buf)
		err error
	)

	if err = enc.Encode(&v); err != nil {
		return err
	}

	output.Stdoutf("[encoder]", "\n%s\n", buf.String())

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: yaml_encoder

	   [Name] "yaml_encoder"
	   [encoder]
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
	           minor: 0



	   [Elapsed] 114.52µs
	*/

}

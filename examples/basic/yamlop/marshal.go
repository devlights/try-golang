package yamlop

import (
	"github.com/devlights/gomy/output"
	"github.com/devlights/try-golang/examples/basic/yamlop/types"
	"gopkg.in/yaml.v3"
)

// Marshal は、yaml.Marshal() を利用したサンプルです.
func Marshal() error {
	var (
		v = types.YamlData{
			Languages: []types.Language{
				{Name: "golang", PrintFn: "fmt.Println", Version: types.Version{Major: 1, Minor: 16}},
				{Name: "java", PrintFn: "System.out.println", Version: types.Version{Major: 16, Minor: 0}},
			},
		}
	)

	var (
		buf []byte
		err error
	)

	if buf, err = yaml.Marshal(&v); err != nil {
		return err
	}

	output.Stdoutf("[yaml]", "\n%s\n", string(buf))

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: yaml_marshal

	   [Name] "yaml_marshal"
	   [yaml]
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



	   [Elapsed] 211µs
	*/

}

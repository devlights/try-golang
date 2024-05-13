package yamlop

import (
	"github.com/devlights/gomy/output"
	"github.com/devlights/try-golang/examples/basic/yamlop/types"
	"gopkg.in/yaml.v3"
)

func Unmarshal() error {
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
		v   types.YamlData
		err error
	)

	if err = yaml.Unmarshal([]byte(yamlStr), &v); err != nil {
		return err
	}

	output.Stdoutf("[unmarshal]", "%v\n", v)

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: yaml_unmarshal

	   [Name] "yaml_unmarshal"
	   [unmarshal]          {[{golang fmt.Println {1 16}} {java System.out.println {16 0}}]}


	   [Elapsed] 177.63µs
	*/

}

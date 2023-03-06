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
}

package yamlop

import (
	"github.com/devlights/gomy/output"
	"github.com/devlights/try-golang/examples/basic/yamlop/types"
	"gopkg.in/yaml.v2"
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
}

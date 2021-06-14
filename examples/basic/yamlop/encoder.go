package yamlop

import (
	"bytes"

	"github.com/devlights/gomy/output"
	"github.com/devlights/try-golang/examples/basic/yamlop/types"
	"gopkg.in/yaml.v2"
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
}

package yamlop

import (
	"github.com/devlights/gomy/output"
	"gopkg.in/yaml.v2"
)

// Marshal は、yaml.Marshal() を利用したサンプルです.
func Marshal() error {
	type (
		Version struct {
			Major int `yaml:"major"`
			Minor int `yaml:"minor"`
		}

		Language struct {
			Name    string  `yaml:"name"`
			PrintFn string  `yaml:"printfn"`
			Version Version `yaml:"version"`
		}

		YamlData struct {
			Languages []Language `yaml:"languages"`
		}
	)

	var (
		v = YamlData{
			Languages: []Language{
				{Name: "golang", PrintFn: "fmt.Println", Version: Version{Major: 1, Minor: 16}},
				{Name: "java", PrintFn: "System.out.println", Version: Version{Major: 16, Minor: 0}},
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

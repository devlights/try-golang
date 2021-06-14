package xmlop

import (
	"encoding/xml"

	"github.com/devlights/gomy/output"
	"github.com/devlights/try-golang/examples/basic/xmlop/types"
)

// MarshalIndent は、xml.MarshalIndent() を使ったサンプルです.
func MarshalIndent() error {
	var (
		v = types.XmlData{
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

	if buf, err = xml.MarshalIndent(&v, "", "    "); err != nil {
		return err
	}

	output.Stdoutf("[marshal]", "\n%s\n", string(buf))

	return nil
}

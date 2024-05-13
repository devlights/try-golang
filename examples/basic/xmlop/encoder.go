package xmlop

import (
	"bytes"
	"encoding/xml"

	"github.com/devlights/gomy/output"
	"github.com/devlights/try-golang/examples/basic/xmlop/types"
)

// Encoder は、xml.Encoder を使ったサンプルです.
func Encoder() error {
	var (
		v = types.XmlData{
			Languages: []types.Language{
				{Name: "golang", PrintFn: "fmt.Println", Version: types.Version{Major: 1, Minor: 16}},
				{Name: "java", PrintFn: "System.out.println", Version: types.Version{Major: 16, Minor: 0}},
			},
		}
	)

	var (
		buf = new(bytes.Buffer)
		enc = xml.NewEncoder(buf)
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

	   ENTER EXAMPLE NAME: xml_encoder

	   [Name] "xml_encoder"
	   [encoder]
	   <data><languages name="golang"><printfn>fmt.Println</printfn><version major="1" minor="16"></version></languages><languages name="java"><printfn>System.out.println</printfn><version major="16" minor="0"></version></languages></data>


	   [Elapsed] 80.37µs
	*/

}

package xmlop

import (
	"bytes"
	"encoding/xml"

	"github.com/devlights/gomy/output"
	"github.com/devlights/try-golang/examples/basic/xmlop/types"
)

// Unmarshal は、xml.Unmarshal() を使ったサンプルです.
func Unmarshal() error {
	const (
		xmlStr = `
<data>
	<languages name="golang">
		<printfn>fmt.Println</printfn>
		<version major="1" minor="16"></version>
	</languages>
	<languages name="java">
		<printfn>System.out.println</printfn>
		<version major="16" minor="0"></version>
	</languages>
</data>`
	)

	var (
		buf = bytes.NewBufferString(xmlStr)
		v   types.XmlData
		err error
	)

	if err = xml.Unmarshal(buf.Bytes(), &v); err != nil {
		return err
	}

	output.Stdoutf("[unmarshal]", "%v\n", v)

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: xml_unmarshal

	   [Name] "xml_unmarshal"
	   [unmarshal]          {{ data} [{golang fmt.Println {1 16}} {java System.out.println {16 0}}]}


	   [Elapsed] 96.249µs
	*/

}

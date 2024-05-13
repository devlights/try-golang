package xmlop

import (
	"encoding/xml"

	"github.com/devlights/gomy/output"
)

// NestedElements -- 入れ子になっている要素を取得するサンプルです。
//
// REFERENCES:
//   - https://pkg.go.dev/encoding/xml@latest#Unmarshal
func NestedElements() error {
	const (
		xmlStr = `
<data>
	<items>
		<item>
			<value>hello</value>
		</item>
		<item>
			<value>world</value>
		</item>
	</items>
</data>`
	)

	type xmldata struct {
		XMLName xml.Name `xml:"data"`             // ルート要素
		Values  []string `xml:"items>item>value"` // 入れ子の要素
	}

	var (
		data xmldata
		err  error
	)

	err = xml.Unmarshal([]byte(xmlStr), &data)
	if err != nil {
		return err
	}

	for _, v := range data.Values {
		output.Stdoutf("[result]", "%v\n", v)
	}

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: xml_nested_elements

	   [Name] "xml_nested_elements"
	   [result]             hello
	   [result]             world


	   [Elapsed] 54.25µs
	*/

}

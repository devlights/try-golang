package txttmpl

import (
	"bytes"
	"text/template"

	"github.com/devlights/gomy/output"
)

// Gt -- text/template の テンプレート仕様 における gt (greater than) のサンプルです.
//
// # REFERENCES
//   - https://pkg.go.dev/text/template@go1.17.8
func Gt() error {
	var (
		tmpls = []string{
			// 10 > 10
			`{{ $v := 10 }} {{- if gt $v 10 -}} helloworld {{- end}}`,
			// 10 > 11
			`{{ $v := 10 }} {{- if gt $v 11 -}} helloworld {{- end}}`,
			// 10 > 9
			`{{ $v := 10 }} {{- if gt $v 9 -}} helloworld {{- end}}`,
		}
	)

	for _, t := range tmpls {
		var (
			tmpl *template.Template
			buf  bytes.Buffer
			err  error
		)

		output.Stdoutl("[template]", t)

		tmpl, err = template.New("Gt").Parse(t)
		if err != nil {
			return err
		}

		err = tmpl.Execute(&buf, nil)
		if err != nil {
			return err
		}

		output.Stdoutf("[tmpl]", "%q\n", buf.String())
		output.StdoutHr()
	}

	return nil
}

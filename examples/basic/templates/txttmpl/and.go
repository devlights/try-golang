package txttmpl

import (
	"bytes"
	"text/template"

	"github.com/devlights/gomy/output"
)

// And -- text/template の テンプレート仕様 における and のサンプルです.
//
// # REFERENCES
//   - https://pkg.go.dev/text/template@go1.17.8
func And() error {
	var (
		tmpls = []string{
			`{{ if and (eq .Hello "hello") (eq .World "world") }} hello world {{end}}`,
			`{{ if and (eq .Hello "world") (eq .World "hello") }} hello world {{end}}`,
		}
	)

	for _, t := range tmpls {
		var (
			tmpl *template.Template
			buf  bytes.Buffer
			err  error
		)

		output.Stdoutl("[template]", t)

		tmpl, err = template.New("And").Parse(t)
		if err != nil {
			return err
		}

		err = tmpl.Execute(&buf, struct{ Hello, World string }{"hello", "world"})
		if err != nil {
			return err
		}

		output.Stdoutf("[tmpl]", "%q\n", buf.String())
		output.StdoutHr()
	}

	return nil
}

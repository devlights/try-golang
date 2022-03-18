package txttmpl

import (
	"bytes"
	"text/template"

	"github.com/devlights/gomy/output"
)

// Pipe -- text/template の テンプレート仕様 における pipe のサンプルです.
//
// # REFERENCES
//   - https://pkg.go.dev/text/template@latest
func Pipe() error {
	var (
		tmpls = []string{
			`{{ "hello" | printf "%s >world<" | html | printf "%q" }}`,
		}
	)

	for _, t := range tmpls {
		var (
			tmpl *template.Template
			buf  bytes.Buffer
			err  error
		)

		output.Stdoutl("[template]", t)

		tmpl, err = template.New("Pipe").Parse(t)
		if err != nil {
			return err
		}

		err = tmpl.Execute(&buf, nil)
		if err != nil {
			return err
		}

		output.Stdoutf("[tmpl]", "%s\n", buf.String())
		output.StdoutHr()
	}

	return nil
}

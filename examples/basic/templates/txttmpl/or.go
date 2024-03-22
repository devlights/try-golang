package txttmpl

import (
	"bytes"
	"text/template"

	"github.com/devlights/gomy/output"
)

// Or -- text/template の テンプレート仕様 における or のサンプルです.
//
// # REFERENCES
//   - https://pkg.go.dev/text/template@latest
func Or() error {
	var (
		tmpls = []string{
			`{{ if or (eq .Hello "hello") (eq .World "world") }} hello world {{end}}`,
			`{{ if or (eq .Hello "world") (eq .World "world") }} hello world {{end}}`,
		}
	)

	for _, t := range tmpls {
		var (
			tmpl *template.Template
			buf  bytes.Buffer
			err  error
		)

		output.Stdoutl("[template]", t)

		tmpl, err = template.New("Or").Parse(t)
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

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: templates_text_tmpl_or

	   [Name] "templates_text_tmpl_or"
	   [template]           {{ if or (eq .Hello "hello") (eq .World "world") }} hello world {{end}}
	   [tmpl]               " hello world "
	   --------------------------------------------------
	   [template]           {{ if or (eq .Hello "world") (eq .World "world") }} hello world {{end}}
	   [tmpl]               " hello world "
	   --------------------------------------------------


	   [Elapsed] 123.56µs
	*/

}

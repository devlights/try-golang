package txttmpl

import (
	"bytes"
	"text/template"

	"github.com/devlights/gomy/output"
)

// Not -- text/template の テンプレート仕様 における not のサンプルです.
//
// # REFERENCES
//   - https://pkg.go.dev/text/template@latest
func Not() error {
	var (
		tmpls = []string{
			`{{ if not (eq .Hello "world") }} hello world {{end}}`,
			`{{ if not (gt (len .Hello) 5) }} hello world {{end}}`,
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

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: templates_text_tmpl_not

	   [Name] "templates_text_tmpl_not"
	   [template]           {{ if not (eq .Hello "world") }} hello world {{end}}
	   [tmpl]               " hello world "
	   --------------------------------------------------
	   [template]           {{ if not (gt (len .Hello) 5) }} hello world {{end}}
	   [tmpl]               " hello world "
	   --------------------------------------------------


	   [Elapsed] 195.13µs
	*/

}

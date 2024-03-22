package txttmpl

import (
	"bytes"
	"text/template"

	"github.com/devlights/gomy/output"
)

// And -- text/template の テンプレート仕様 における and のサンプルです.
//
// # REFERENCES
//   - https://pkg.go.dev/text/template@latest
func With() error {
	var (
		tmpls = []string{
			`{{ with "hello" }}{{ printf "%s world" . }}{{end}}`,
			`{{ with $v := "hello" }}{{ printf "%s world" $v }}{{end}}`,
		}
	)

	for _, t := range tmpls {
		var (
			tmpl *template.Template
			buf  bytes.Buffer
			err  error
		)

		output.Stdoutl("[template]", t)

		tmpl, err = template.New("With").Parse(t)
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

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: templates_text_tmpl_with

	   [Name] "templates_text_tmpl_with"
	   [template]           {{ with "hello" }}{{ printf "%s world" . }}{{end}}
	   [tmpl]               hello world
	   --------------------------------------------------
	   [template]           {{ with $v := "hello" }}{{ printf "%s world" $v }}{{end}}
	   [tmpl]               hello world
	   --------------------------------------------------


	   [Elapsed] 123.6µs
	*/

}

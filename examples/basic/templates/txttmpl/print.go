package txttmpl

import (
	"bytes"
	"text/template"

	"github.com/devlights/gomy/output"
)

// Print -- text/template の テンプレート仕様 における print 関数 のサンプルです.
//
// # REFERENCES
//   - https://pkg.go.dev/text/template@latest
func Print() error {
	var (
		tmpls = []string{
			`{{ print . }}`,
		}
	)

	for _, t := range tmpls {
		var (
			tmpl *template.Template
			buf  bytes.Buffer
			err  error
		)

		output.Stdoutl("[template]", t)

		tmpl, err = template.New("Print").Parse(t)
		if err != nil {
			return err
		}

		err = tmpl.Execute(&buf, "hello\tworld")
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

	   ENTER EXAMPLE NAME: templates_text_tmpl_print

	   [Name] "templates_text_tmpl_print"
	   [template]           {{ print . }}
	   [tmpl]               hello      world
	   --------------------------------------------------


	   [Elapsed] 76.54µs
	*/

}

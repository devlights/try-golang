package txttmpl

import (
	"bytes"
	"text/template"

	"github.com/devlights/gomy/output"
)

// Urlquery -- text/template の テンプレート仕様 における urlquery 関数 のサンプルです.
//
// # REFERENCES
//   - https://pkg.go.dev/text/template@latest
func Urlquery() error {
	var (
		tmpls = []string{
			`{{ urlquery . }}`,
		}
	)

	for _, t := range tmpls {
		var (
			tmpl *template.Template
			buf  bytes.Buffer
			err  error
		)

		output.Stdoutl("[template]", t)

		tmpl, err = template.New("Urlquery").Parse(t)
		if err != nil {
			return err
		}

		err = tmpl.Execute(&buf, "https://jsonplaceholder.typicode.com/comments?postId=1")
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

	   ENTER EXAMPLE NAME: templates_text_tmpl_urlquery

	   [Name] "templates_text_tmpl_urlquery"
	   [template]           {{ urlquery . }}
	   [tmpl]               "https%3A%2F%2Fjsonplaceholder.typicode.com%2Fcomments%3FpostId%3D1"
	   --------------------------------------------------


	   [Elapsed] 139.89µs
	*/

}

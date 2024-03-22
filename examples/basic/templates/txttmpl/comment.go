package txttmpl

import (
	"bytes"
	"text/template"

	"github.com/devlights/gomy/output"
)

// Comment -- text/template の テンプレート仕様 におけるコメントのサンプルです.
//
// # REFERENCES
//   - https://pkg.go.dev/text/template@latest
func Comment() error {
	const (
		t = `this is {{/* comment */}}     {{- /* another comment */ -}}.`
	)

	var (
		tmpl *template.Template
		buf  bytes.Buffer
		err  error
	)

	tmpl, err = template.New("Comment").Parse(t)
	if err != nil {
		return err
	}

	err = tmpl.Execute(&buf, nil)
	if err != nil {
		return err
	}

	output.Stdoutl("[comment]", buf.String())

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: templates_text_tmpl_comment

	   [Name] "templates_text_tmpl_comment"
	   [comment]            this is .


	   [Elapsed] 90.53µs
	*/

}

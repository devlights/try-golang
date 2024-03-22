package txttmpl

import (
	"bytes"
	"text/template"

	"github.com/devlights/gomy/output"
)

// Define -- text/template の テンプレート仕様 における define (独自テンプレートの定義) のサンプルです.
//
// # REFERENCES
//   - https://pkg.go.dev/text/template@latest
//   - https://stackoverflow.com/questions/37298532/variable-inside-define-template-in-golang
func Define() error {
	var (
		tmpls = []string{
			`{{ define "mytmpl" }}||hello world||{{ end }} {{ template "mytmpl" }}`,
			`{{ define "mytmpl" }}||hello world ({{.V}})||{{ end }} {{ template "mytmpl" . }}`,
		}
	)

	for _, t := range tmpls {
		var (
			tmpl *template.Template
			buf  bytes.Buffer
			err  error
		)

		output.Stdoutl("[template]", t)

		tmpl, err = template.New("Define").Parse(t)
		if err != nil {
			return err
		}

		err = tmpl.Execute(&buf, struct{ V int }{100})
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

	   ENTER EXAMPLE NAME: templates_text_tmpl_define

	   [Name] "templates_text_tmpl_define"
	   [template]           {{ define "mytmpl" }}||hello world||{{ end }} {{ template "mytmpl" }}
	   [tmpl]                ||hello world||
	   --------------------------------------------------
	   [template]           {{ define "mytmpl" }}||hello world ({{.V}})||{{ end }} {{ template "mytmpl" . }}
	   [tmpl]                ||hello world (100)||
	   --------------------------------------------------


	   [Elapsed] 131.48µs
	*/

}

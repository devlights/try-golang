package txttmpl

import (
	"bytes"
	"text/template"

	"github.com/devlights/gomy/output"
)

// If -- text/template の テンプレート仕様 におけるifのサンプルです.
//
// # REFERENCES
//   - https://pkg.go.dev/text/template@latest
//   - https://ema-hiro.hatenablog.com/entry/20170729/1501320887
//   - https://stackoverflow.com/questions/33282061/go-templates-with-eq-and-index
//   - https://stackoverflow.com/questions/31101729/compare-strings-in-templates
func If() error {
	const (
		t1 = `{{if .}} hello world {{end}}`
		t2 = `{{if . -}} hello world {{- end}}`
	)

	for _, t := range []string{t1, t2} {
		var (
			tmpl *template.Template
			buf  bytes.Buffer
			err  error
		)

		output.Stdoutl("[template]", t)

		tmpl, err = template.New("If").Parse(t)
		if err != nil {
			return err
		}

		err = tmpl.Execute(&buf, true)
		if err != nil {
			return err
		}
		output.Stdoutf("[true ]", "%q\n", buf.String())

		buf.Reset()
		err = tmpl.Execute(&buf, false)
		if err != nil {
			return err
		}
		output.Stdoutf("[false]", "%q\n", buf.String())
		output.StdoutHr()
	}

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: templates_text_tmpl_if

	   [Name] "templates_text_tmpl_if"
	   [template]           {{if .}} hello world {{end}}
	   [true ]              " hello world "
	   [false]              ""
	   --------------------------------------------------
	   [template]           {{if . -}} hello world {{- end}}
	   [true ]              "hello world"
	   [false]              ""
	   --------------------------------------------------


	   [Elapsed] 156.51µs
	*/

}

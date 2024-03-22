package txttmpl

import (
	"bytes"
	"text/template"

	"github.com/devlights/gomy/output"
)

// Range -- text/template の テンプレート仕様 におけるrangeのサンプルです.
//
// # REFERENCES
//   - https://pkg.go.dev/text/template@latest
//   - https://stackoverflow.com/questions/54156119/range-over-string-slice-in-golang-template
func Range() error {
	const (
		t1 = `{{range .Items}} {{.}} {{end}}`
		t2 = `{{range $elem := .Items}} {{$elem}} {{end}}`
		t3 = `{{range $i,$elem := .Items}} {{$i}}:{{$elem}} {{end}}`
	)
	type (
		D struct {
			Items []string
		}
	)

	for _, t := range []string{t1, t2, t3} {
		var (
			tmpl *template.Template
			buf  bytes.Buffer
			err  error
		)

		output.Stdoutl("[template]", t)

		tmpl, err = template.New("Range").Parse(t)
		if err != nil {
			return err
		}

		err = tmpl.Execute(&buf, D{[]string{"hello", "world"}})
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

	   ENTER EXAMPLE NAME: templates_text_tmpl_range

	   [Name] "templates_text_tmpl_range"
	   [template]           {{range .Items}} {{.}} {{end}}
	   [tmpl]               " hello  world "
	   --------------------------------------------------
	   [template]           {{range $elem := .Items}} {{$elem}} {{end}}
	   [tmpl]               " hello  world "
	   --------------------------------------------------
	   [template]           {{range $i,$elem := .Items}} {{$i}}:{{$elem}} {{end}}
	   [tmpl]               " 0:hello  1:world "
	   --------------------------------------------------


	   [Elapsed] 154.03µs
	*/

}

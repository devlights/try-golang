package txttmpl

import (
	"bytes"
	"text/template"

	"github.com/devlights/gomy/output"
)

// Else -- text/template の テンプレート仕様 におけるelseのサンプルです.
//
// # REFERENCES
//   - https://pkg.go.dev/text/template@latest
//   - https://ema-hiro.hatenablog.com/entry/20170729/1501320887
//   - https://stackoverflow.com/questions/33282061/go-templates-with-eq-and-index
//   - https://stackoverflow.com/questions/31101729/compare-strings-in-templates
func Else() error {
	const (
		t1 = `{{if eq .Message "hello" -}} world {{- else if eq .Message "world" -}} hello {{- else -}} helloworld {{- end}}`
	)
	type (
		D struct {
			Message string
		}
	)

	for _, t := range []string{t1} {
		var (
			tmpl *template.Template
			buf  bytes.Buffer
			err  error
		)

		tmpl, err = template.New("ElseIf").Parse(t)
		if err != nil {
			return err
		}

		err = tmpl.Execute(&buf, D{"golang"})
		if err != nil {
			return err
		}

		output.Stdoutl("[tmpl]", buf.String())
	}

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: templates_text_tmpl_else

	   [Name] "templates_text_tmpl_else"
	   [tmpl]               helloworld


	   [Elapsed] 192.2µs
	*/

}

package htmltmpl

import (
	"bytes"
	"html/template"

	"github.com/devlights/gomy/output"
)

func Escape() error {
	var (
		tmpls = []string{
			`{{ . }}`,
		}
	)

	for _, t := range tmpls {
		var (
			tmpl *template.Template
			buf  bytes.Buffer
			err  error
		)

		output.Stdoutl("[template]", t)

		tmpl, err = template.New("Escape").Parse(t)
		if err != nil {
			return err
		}

		err = tmpl.Execute(&buf, "<script>alert('hello world')</script>")
		if err != nil {
			return err
		}

		output.Stdoutf("[tmpl]", "%s\n", buf.String())
		output.StdoutHr()
	}

	return nil
}

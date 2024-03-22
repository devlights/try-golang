package txttmpl

import (
	"bytes"
	"text/template"

	"github.com/devlights/gomy/output"
)

// Structure -- テンプレートに差し込む構造体についてのサンプルです.
//
// # REFERENCES
//   - https://pkg.go.dev/text/template@latest
func Structure() error {
	const (
		t = `{{ .Header }} {{ .Footer }}`
	)

	type (
		D struct {
			Header, Footer string
		}
	)

	var (
		tmpl *template.Template
		buf  bytes.Buffer
		err  error
		data = D{"hello", "world"}
	)

	tmpl, err = template.New("Structure").Parse(t)
	if err != nil {
		return err
	}

	err = tmpl.Execute(&buf, data)
	if err != nil {
		return err
	}

	output.Stdoutl("[template ]", t)
	output.Stdoutl("[tmpl-exec]", buf.String())

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: templates_text_tmpl_structure

	   [Name] "templates_text_tmpl_structure"
	   [template ]          {{ .Header }} {{ .Footer }}
	   [tmpl-exec]          hello world


	   [Elapsed] 108.991µs
	*/

}

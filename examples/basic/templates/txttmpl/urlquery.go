package txttmpl

import (
	"bytes"
	"text/template"

	"github.com/devlights/gomy/output"
)

// Urlquery -- text/template の テンプレート仕様 における urlquery 関数 のサンプルです.
//
// # REFERENCES
//   - https://pkg.go.dev/text/template@go1.17.8
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

		err = tmpl.Execute(&buf, "comments?postId=1")
		if err != nil {
			return err
		}

		output.Stdoutf("[tmpl]", "%q\n", buf.String())
		output.StdoutHr()
	}

	return nil
}

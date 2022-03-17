package txttmpl

import (
	"bytes"
	"text/template"

	"github.com/devlights/gomy/output"
)

// Html -- text/template の テンプレート仕様 における html のサンプルです.
//
// # REFERENCES
//   - https://pkg.go.dev/text/template@latest
func Html() error {
	var (
		tmpls = []string{
			`元: {{ printf "%q\t" . }}後: {{ printf "%q" (html .) }}`,
		}
	)

	for _, t := range tmpls {
		var (
			tmpl *template.Template
			buf  bytes.Buffer
			err  error
		)

		output.Stdoutl("[template]", t)

		tmpl, err = template.New("Html").Parse(t)
		if err != nil {
			return err
		}

		err = tmpl.Execute(&buf, "<>'\"& ")
		if err != nil {
			return err
		}

		output.Stdoutf("[tmpl]", "%s\n", buf.String())
		output.StdoutHr()
	}

	return nil
}

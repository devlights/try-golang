package parsefuncs

import (
	"bytes"
	"path/filepath"
	"text/template"

	"github.com/devlights/gomy/output"
)

// ParseGlob -- template.ParseGlob() のサンプルです。
//
// REFERENCES
//   - https://pkg.go.dev/text/template@latest#ParseGlob
func ParseGlob() error {
	const (
		dir = "examples/basic/templates/txttmpl/parsefuncs/tmpls"
		pat = "*.tmpl"
	)

	var (
		pattern = filepath.Join(dir, pat)
		tmpl    *template.Template
		err     error
	)

	output.Stdoutl("[pattern]", pattern)
	output.StdoutHr()

	tmpl, err = template.ParseGlob(pattern)
	if err != nil {
		return err
	}

	for _, t := range tmpl.Templates() {
		var (
			buf bytes.Buffer
		)

		err = t.Execute(&buf, "hello")
		if err != nil {
			return err
		}

		output.Stdoutl("[name       ]", t.Name())
		output.Stdoutl("[parse name ]", t.ParseName)
		output.Stdoutf("[tmpl       ]", "%s\n", buf.String())
		output.StdoutHr()
	}

	return nil
}

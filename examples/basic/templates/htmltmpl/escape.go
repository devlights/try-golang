package htmltmpl

import (
	"bytes"
	"html/template"

	"github.com/devlights/gomy/output"
)

// Escape は、html/template にて適用されるHTMLエスケープについてのサンプルです.
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

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: templates_html_tmpl_escape

	   [Name] "templates_html_tmpl_escape"
	   [template]           {{ . }}
	   [tmpl]               &lt;script&gt;alert(&#39;hello world&#39;)&lt;/script&gt;
	   --------------------------------------------------


	   [Elapsed] 116.43µs
	*/

}

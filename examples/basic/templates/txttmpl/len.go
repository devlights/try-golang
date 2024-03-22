package txttmpl

import (
	"bytes"
	"text/template"

	"github.com/devlights/gomy/output"
)

// Len -- text/template の テンプレート仕様 における len のサンプルです.
//
// # REFERENCES
//   - https://pkg.go.dev/text/template@latest
func Len() error {
	var (
		tmpls = []string{
			`{{ printf "%T:%d\t%T:%d\t%T:%d\t%T:%d" .S (len .S) .M (len .M) .A (len .A) .Str (len .Str) }}`,
		}
	)

	for _, t := range tmpls {
		var (
			tmpl *template.Template
			buf  bytes.Buffer
			err  error
		)

		output.Stdoutl("[template]", t)

		tmpl, err = template.New("Len").Parse(t)
		if err != nil {
			return err
		}

		d := struct {
			S   []string
			M   map[string]string
			A   [1]string
			Str string
		}{
			[]string{"hello", "world"},
			map[string]string{"hello": "world"},
			[1]string{"hello"},
			"helloworld",
		}

		err = tmpl.Execute(&buf, d)
		if err != nil {
			return err
		}

		output.Stdoutf("[tmpl]", "%v\n", buf.String())
		output.StdoutHr()
	}

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: templates_text_tmpl_len

	   [Name] "templates_text_tmpl_len"
	   [template]           {{ printf "%T:%d\t%T:%d\t%T:%d\t%T:%d" .S (len .S) .M (len .M) .A (len .A) .Str (len .Str) }}
	   [tmpl]               []string:2 map[string]string:1     [1]string:1     string:10
	   --------------------------------------------------


	   [Elapsed] 230.36µs
	*/

}

package parsefuncs

import (
	"bytes"
	"os"
	"text/template"

	"github.com/devlights/gomy/output"
)

// ParseFS -- template.ParseFS() のサンプルです。
//
// REFERENCES
//   - https://pkg.go.dev/text/template@latest#ParseFS
//   - https://pkg.go.dev/io/fs@latest
func ParseFS() error {
	const (
		dir = "examples/basic/templates/txttmpl/parsefuncs/tmpls"
	)

	var (
		fSys = os.DirFS(dir)
		tmpl *template.Template
		buf  bytes.Buffer
		err  error
	)

	tmpl, err = template.ParseFS(fSys, "hello.tmpl")
	if err != nil {
		return err
	}

	output.Stdoutl("[name       ]", tmpl.Name())
	output.Stdoutl("[parse name ]", tmpl.ParseName)
	output.Stdoutl("[define tmpl]", tmpl.DefinedTemplates())
	output.StdoutHr()

	err = tmpl.Execute(&buf, "world")
	if err != nil {
		return err
	}

	output.Stdoutf("[tmpl]", "%s\n", buf.String())
	output.StdoutHr()

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: templates_parsefuncs_parsefs

	   [Name] "templates_parsefuncs_parsefs"
	   [name       ]        hello.tmpl
	   [parse name ]        hello.tmpl
	   [define tmpl]        ; defined templates are: "hello.tmpl"
	   --------------------------------------------------
	   [tmpl]               [hello.tmpl] hello world
	   --------------------------------------------------


	   [Elapsed] 100.42µs
	*/

}

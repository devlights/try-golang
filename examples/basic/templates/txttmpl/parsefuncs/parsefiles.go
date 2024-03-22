package parsefuncs

import (
	"bytes"
	"io/fs"
	"os"
	"path/filepath"
	"text/template"

	"github.com/devlights/gomy/output"
)

// ParseFiles -- template.ParseFiles() のサンプルです。
//
// REFERENCES
//   - https://pkg.go.dev/text/template@latest#ParseFiles
func ParseFiles() error {
	const (
		dir = "examples/basic/templates/txttmpl/parsefuncs/tmpls"
	)

	var (
		fSys  = os.DirFS(dir)
		files []string
		err   error
	)

	files, err = fs.Glob(fSys, "*.tmpl")
	if err != nil {
		return err
	}

	for i, f := range files {
		files[i] = filepath.Join(dir, f)
	}

	var (
		tmpl *template.Template
		buf  bytes.Buffer
	)

	tmpl, err = template.ParseFiles(files...)
	if err != nil {
		return err
	}

	output.Stdoutl("[name       ]", tmpl.Name())
	output.Stdoutl("[parse name ]", tmpl.ParseName)
	output.Stdoutl("[define tmpl]", tmpl.DefinedTemplates())
	output.StdoutHr()

	//
	// 複数のテンプレートが存在する場合 Execute の呼び出しは最初のテンプレートが処理される.
	// 最初のテンプレートとは tmpl.ParseName が返すテンプレートの値.
	//
	err = tmpl.Execute(&buf, "world")
	if err != nil {
		return err
	}

	output.Stdoutf("[tmpl - Execute        ]", "%s\n", buf.String())
	output.StdoutHr()

	//
	// 特定のテンプレートを処理したい場合は、ExecuteTemplate を使う
	//
	for _, name := range []string{"hello.tmpl", "world.tmpl"} {
		buf.Reset()

		err = tmpl.ExecuteTemplate(&buf, name, "hello")
		if err != nil {
			return err
		}

		output.Stdoutf("[tmpl - ExecuteTemplate]", "%s\n", buf.String())
		output.StdoutHr()
	}

	//
	// または t.Templates() で各テンプレートが取得できるので、それで処理することもできる
	//
	for _, t := range tmpl.Templates() {
		buf.Reset()

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

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: templates_parsefuncs_parsefiles

	   [Name] "templates_parsefuncs_parsefiles"
	   [name       ]        hello.tmpl
	   [parse name ]        hello.tmpl
	   [define tmpl]        ; defined templates are: "world.tmpl", "hello.tmpl"
	   --------------------------------------------------
	   [tmpl - Execute        ] [hello.tmpl] hello world
	   --------------------------------------------------
	   [tmpl - ExecuteTemplate] [hello.tmpl] hello hello
	   --------------------------------------------------
	   [tmpl - ExecuteTemplate] [world.tmpl] hello world
	   --------------------------------------------------
	   [name       ]        hello.tmpl
	   [parse name ]        hello.tmpl
	   [tmpl       ]        [hello.tmpl] hello hello
	   --------------------------------------------------
	   [name       ]        world.tmpl
	   [parse name ]        world.tmpl
	   [tmpl       ]        [world.tmpl] hello world
	   --------------------------------------------------


	   [Elapsed] 313.69µs
	*/

}

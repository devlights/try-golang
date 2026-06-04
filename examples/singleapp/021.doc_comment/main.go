package main

import (
	"flag"
	"go/doc/comment"
	"log"
	"os"
)

const (
	sample = `
# Title

helloworld.

[path/filepath] package.

[io.Reader] interface.

code block
        package main
        func main(){}

list
        - a
        - b

[path/filepath]: https://pkg.go.dev/path/filepath@go1.23.1
[io.Reader]: https://pkg.go.dev/io@go1.23.1#Reader
`
)

type (
	Args struct {
		markdown bool
		html     bool
		comment  bool
		text     bool
	}
)

var (
	args Args
)

func init() {
	flag.BoolVar(&args.markdown, "md", false, "output markdown")
	flag.BoolVar(&args.html, "html", false, "output html")
	flag.BoolVar(&args.comment, "comment", false, "output go comment")
	flag.BoolVar(&args.text, "text", false, "output text")
}

func main() {
	log.SetFlags(0)
	flag.Parse()

	if err := run(); err != nil {
		log.Panic(err)
	}
}

func run() error {
	var (
		parser  comment.Parser
		printer comment.Printer
		doc     *comment.Doc
		buf     []byte
	)
	doc = parser.Parse(sample)
	printer.TextPrefix = "// "

	switch {
	case args.markdown:
		buf = printer.Markdown(doc)
	case args.html:
		buf = printer.HTML(doc)
	case args.comment:
		buf = printer.Comment(doc)
	case args.text:
		buf = printer.Text(doc)
	default:
		flag.PrintDefaults()
		return nil
	}

	if _, err := os.Stdout.Write(buf); err != nil {
		return err
	}

	return nil
}

package main

import (
	"bytes"
	"flag"
	"log"
	"os"
)

var (
	BOM = []byte{0xEF, 0xBB, 0xBF}
)

var (
	files []string
	rmBOM bool
)

func init() {
	log.SetFlags(0)

	flag.BoolVar(&rmBOM, "d", false, "BOM removal mode")
	flag.Usage = func() {
		log.Println("Usage: bom (-d) file-path...")
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()

	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(-1)
	}

	files = make([]string, len(flag.Args()))
	copy(files, flag.Args())

	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	for _, fpath := range files {
		if err := process(fpath); err != nil {
			return err
		}
	}

	return nil
}

func process(fpath string) error {
	fi, err := os.Stat(fpath)
	if err != nil {
		return err
	}

	data, err := os.ReadFile(fpath)
	if err != nil {
		return err
	}

	if bytes.HasPrefix(data, BOM) {
		if rmBOM {
			if err := os.WriteFile(fpath, data[len(BOM):], fi.Mode()); err != nil {
				return err
			}
		}

		return nil
	}

	if err := os.WriteFile(fpath, append(BOM, data...), fi.Mode()); err != nil {
		return err
	}

	return nil

	/*
	   $ task
	   task: [run] echo helloworld > test.txt
	   task: [run] nkf -g test.txt; hexdump test.txt
	   ASCII
	   0000000 6568 6c6c 776f 726f 646c 000a
	   000000b
	   task: [run] go run main.go test.txt
	   task: [run] nkf -g test.txt; hexdump test.txt
	   UTF-8
	   0000000 bbef 68bf 6c65 6f6c 6f77 6c72 0a64
	   000000e
	   task: [run] go run main.go -d test.txt
	   task: [run] nkf -g test.txt; hexdump test.txt
	   ASCII
	   0000000 6568 6c6c 776f 726f 646c 000a
	   000000b
	   task: [run] rm -f test.txt
	*/

}

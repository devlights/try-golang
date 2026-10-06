package main

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	//
	// Go 1.27 で strings.CutLast 関数が追加された。
	// 元々存在している strings.Cut の末尾版
	//
	// - https://pkg.go.dev/strings@go1.27.1#Cut
	// - https://pkg.go.dev/strings@go1.27.1#CutLast
	// - https://pkg.go.dev/strings@go1.27.1#CutPrefix
	// - https://pkg.go.dev/strings@go1.27.1#CutSuffix
	//
	var (
		w = tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.Debug)
	)
	defer w.Flush()

	var (
		s      = "helloworld"
		before string
		after  string
		found  bool
	)
	fmt.Printf("target: %s\n\n", s)

	// strings.Cut      : 先頭から進み、ヒットした箇所でカット
	if before, after, found = strings.Cut(s, "o"); found {
		w.Write(fmt.Appendf(nil, "cut('o')\t%s\t%s\n", before, after))
	}

	// strings.CutLast  : 末尾から進み、ヒットした箇所でカット
	if before, after, found = strings.CutLast(s, "o"); found {
		w.Write(fmt.Appendf(nil, "cutlast('o')\t%s\t%s\n", before, after))
	}

	// strings.CutPrefix: 指定プレフィックスでカット
	if after, found = strings.CutPrefix(s, "hello"); found {
		w.Write(fmt.Appendf(nil, "cutprefix('ll')\t \t%s\n", after))
	}

	// strings.CutSuffix: 指定サフィックスでカット
	if before, found = strings.CutSuffix(s, "world"); found {
		w.Write(fmt.Appendf(nil, "cutsuffix('ll')\t%s\t\n", before))
	}

	return nil
}

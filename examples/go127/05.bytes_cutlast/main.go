package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"text/tabwriter"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	//
	// Go 1.27 で bytes.CutLast 関数が追加された。
	// 元々存在している bytes.Cut の末尾版
	//
	// - https://pkg.go.dev/bytes@go1.27.1#Cut
	// - https://pkg.go.dev/bytes@go1.27.1#CutLast
	// - https://pkg.go.dev/bytes@go1.27.1#CutPrefix
	// - https://pkg.go.dev/bytes@go1.27.1#CutSuffix
	//
	var (
		w = tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.Debug)
	)
	defer w.Flush()

	var (
		b      = make([]byte, 4)
		before []byte
		after  []byte
		found  bool
	)
	binary.BigEndian.PutUint32(b, (uint32)(0xdeadbeef))
	fmt.Printf("target: 0x%x\n\n", b)

	// bytes.Cut      : 先頭から進み、ヒットした箇所でカット
	var (
		sep = make([]byte, 1)
	)
	sep[0] = (byte)(0xbe)

	if before, after, found = bytes.Cut(b, sep); found {
		w.Write(fmt.Appendf(nil, "cut(0xbe)\t%x\t%x\n", before, after))
	}

	// bytes.CutLast  : 末尾から進み、ヒットした箇所でカット
	if before, after, found = bytes.CutLast(b, sep); found {
		w.Write(fmt.Appendf(nil, "cutlast(0xbe)\t%x\t%x\n", before, after))
	}

	// bytes.CutPrefix: 指定プレフィックスでカット
	var (
		prefix = make([]byte, 2)
	)
	binary.BigEndian.PutUint16(prefix, (uint16)(0xdead))

	if after, found = bytes.CutPrefix(b, prefix); found {
		w.Write(fmt.Appendf(nil, "cutprefix(0xdead)\t \t%x\n", after))
	}

	// bytes.CutSuffix: 指定サフィックスでカット
	var (
		suffix = make([]byte, 2)
	)
	binary.BigEndian.PutUint16(suffix, (uint16)(0xbeef))

	if before, found = bytes.CutSuffix(b, suffix); found {
		w.Write(fmt.Appendf(nil, "cutsuffix(0xbeef)\t%x\t\n", before))
	}

	return nil
}

package main

import (
	"flag"
	"fmt"
	"runtime"
	"weak"
)

type (
	Args struct {
		UseWeakRef bool
	}

	MemStats struct {
		// ヒープに割り当てられたメモリの合計 (bytes)
		HeapAlloc uint64
	}
)

const (
	_          = iota
	KB float64 = 1 << (10 * iota)
	MB
	GB
)

var (
	args Args
)

func init() {
	flag.BoolVar(&args.UseWeakRef, "weakref", false, "use weak")
}

func main() {
	flag.Parse()

	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	const (
		oneMib  = 1 << 20     // 1Mib
		bufSize = oneMib << 5 // 32Mib
	)

	switch args.UseWeakRef {
	case true:
		printMem("init  ", getMemStats())

		var (
			x                            = make([]byte, bufSize)
			weakRef weak.Pointer[[]byte] = weak.Make(&x) // https://pkg.go.dev/weak@go1.24.1
		)
		printMem("before", getMemStats())
		runtime.GC()
		printMem("after ", getMemStats())

		fmt.Printf("object is nil? ==> %v\n", weakRef.Value() == nil)

		//
		// runtime.KeepAlive() は、この呼び出しの時点まで引数に指定した
		// オブジェクトが生存していることをコンパイラとGCに対して保証させるもの。
		//
		runtime.KeepAlive(&weakRef)
	default:
		printMem("init  ", getMemStats())

		var (
			x         = make([]byte, bufSize)
			strongRef = struct{ b *[]byte }{&x}
		)
		printMem("before", getMemStats())
		runtime.GC()
		printMem("after ", getMemStats())

		fmt.Printf("object is nil? ==> %v\n", strongRef.b == nil)

		runtime.KeepAlive(&strongRef)
	}

	return nil
}

// メモリサイズをより読みやすい形式に変換
func formatBytes(bytes uint64) string {
	switch {
	case bytes >= uint64(GB):
		return fmt.Sprintf("%7.2f GB", float64(bytes)/GB)
	case bytes >= uint64(MB):
		return fmt.Sprintf("%7.2f MB", float64(bytes)/MB)
	case bytes >= uint64(KB):
		return fmt.Sprintf("%7.2f KB", float64(bytes)/KB)
	default:
		return fmt.Sprintf("%d bytes", bytes)
	}
}

func printMem(prefix string, stats MemStats) {
	fmt.Printf("[%s] HeapAlloc:   %s (ヒープメモリ)\n", prefix, formatBytes(stats.HeapAlloc))
}

func getMemStats() MemStats {
	var rtm runtime.MemStats

	runtime.ReadMemStats(&rtm)

	return MemStats{
		HeapAlloc: rtm.HeapAlloc,
	}
}

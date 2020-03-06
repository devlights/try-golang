package mem

import (
	"fmt"
	"github.com/devlights/try-golang/lib/output"
	"runtime"
)

type (
	// Mem は、実行時メモリの情報を出力する振る舞いを定義するインターフェースです.
	Mem interface {
		// 実行時メモリの情報を出力します.
		Print(prefix string)
		setAlloc(enabled bool)
		setHeapAlloc(enabled bool)
		setTotalAlloc(enabled bool)
		setHeapObjects(enabled bool)
		setSys(enabled bool)
		setNumGC(enabled bool)
	}

	runtimeMem struct {
		enableAlloc       bool
		enableHeapAlloc   bool
		enableTotalAlloc  bool
		enableHeapObjects bool
		enableSys         bool
		enableNumGC       bool
	}

	Option func(Mem)
)

// NewMem は、指定された情報を元に Mem を生成します.
//
// デフォルトでは、以下の項目が有効となります.
//   - Alloc
//   - TotalAlloc
//   - NumGC
func NewMem(options ...Option) Mem {
	m := new(runtimeMem)

	defaultOptions := []Option{
		Alloc(true),
		HeapAlloc(false),
		TotalAlloc(true),
		HeapObjects(false),
		Sys(false),
		NumGC(true),
	}

	for _, option := range append(defaultOptions, options...) {
		option(m)
	}

	return m
}

// Alloc の項目を有効にするかどうかを指定します.
//
//noinspection GoUnusedExportedFunction
func Alloc(enabled bool) Option {
	return func(mem Mem) {
		mem.setAlloc(enabled)
	}
}

// HeapAlloc の項目を有効にするかどうかを指定します.
//
//noinspection GoUnusedExportedFunction
func HeapAlloc(enabled bool) Option {
	return func(mem Mem) {
		mem.setHeapAlloc(enabled)
	}
}

// TotalAlloc の項目を有効にするかどうかを指定します.
//
//noinspection GoUnusedExportedFunction
func TotalAlloc(enabled bool) Option {
	return func(mem Mem) {
		mem.setTotalAlloc(enabled)
	}
}

// HeapObjects の項目を有効にするかどうかを指定します.
//
//noinspection GoUnusedExportedFunction
func HeapObjects(enabled bool) Option {
	return func(mem Mem) {
		mem.setHeapObjects(enabled)
	}
}

// Sys の項目を有効にするかどうかを指定します.
//
//noinspection GoUnusedExportedFunction
func Sys(enabled bool) Option {
	return func(mem Mem) {
		mem.setSys(enabled)
	}
}

// NumGC の項目を有効にするかどうかを指定します.
//
//noinspection GoUnusedExportedFunction
func NumGC(enabled bool) Option {
	return func(mem Mem) {
		mem.setNumGC(enabled)
	}
}

func (m *runtimeMem) setAlloc(enabled bool) {
	m.enableAlloc = enabled
}

func (m *runtimeMem) setHeapAlloc(enabled bool) {
	m.enableHeapAlloc = enabled
}

func (m *runtimeMem) setTotalAlloc(enabled bool) {
	m.enableTotalAlloc = enabled
}

func (m *runtimeMem) setHeapObjects(enabled bool) {
	m.enableHeapObjects = enabled
}

func (m *runtimeMem) setSys(enabled bool) {
	m.enableSys = enabled
}

func (m *runtimeMem) setNumGC(enabled bool) {
	m.enableNumGC = enabled
}

// Print は、現在のメモリ量などを出力します.
func (m *runtimeMem) Print(prefix string) {
	var (
		ms runtime.MemStats
	)

	output.Stdoutl(prefix, "----------------------------")
	runtime.ReadMemStats(&ms)

	if m.enableAlloc {
		output.Stdoutl("Alloc", toKb(ms.Alloc))
	}

	if m.enableHeapAlloc {
		output.Stdoutl("HeapAlloc", toKb(ms.HeapAlloc))
	}

	if m.enableTotalAlloc {
		output.Stdoutl("TotalAlloc", toKb(ms.TotalAlloc))
	}

	if m.enableHeapObjects {
		output.Stdoutl("HeapObjects", toKb(ms.HeapObjects))
	}

	if m.enableSys {
		output.Stdoutl("Sys", toKb(ms.Sys))
	}

	if m.enableNumGC {
		output.Stdoutl("NumGC", ms.NumGC)
	}
}

func toKb(bytes uint64) string {
	return fmt.Sprintf("%d KiB", bytes/1024)
}

//noinspection GoUnusedFunction
func toMb(bytes uint64) string {
	return fmt.Sprintf("%d MiB", bytes/1024/1024)
}

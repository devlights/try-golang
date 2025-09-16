//go:build unix

// Go 1.19 で Build Constraints に新たに unix が追加された。
//
//   - https://go.dev/doc/go1.19#go-unix
//
// 実際のマッピングは $(go env GOROOT)/src/cmd/dist/build.go に以下のように定義されている。(Go 1.25)
//
//	var unixOS = map[string]bool{
//		"aix":       true,
//		"android":   true,
//		"darwin":    true,
//		"dragonfly": true,
//		"freebsd":   true,
//		"hurd":      true,
//		"illumos":   true,
//		"ios":       true,
//		"linux":     true,
//		"netbsd":    true,
//		"openbsd":   true,
//		"solaris":   true,
//	}
//
// なので、linux, macOS(darwin) の両方で有効にする場合は unix を指定すれば良い。
//
// # REFERENCES
//   - https://cs.opensource.google/go/go/+/refs/tags/go1.25.1:src/cmd/dist/build.go;l=1070
//   - https://dev.to/emreodabas/quick-guide-go-119-features-1j40
//   - https://go.dev/doc/go1.19#go-unix
package main

import "fmt"

func main() {
	fmt.Println("run on unix")
}

//go:build unix

// Go 1.19 で Build Constraints に新たに unix が追加された。
// GOOS が以下の場合に有効になる。
//
//	ios, linux, android, darwin, dragonfly, freebsd, hurd, illumos, netbsd, aix, openbsd or solaris
//
// # REFERENCES
//   - https://dev.to/emreodabas/quick-guide-go-119-features-1j40
//   - https://go.dev/doc/go1.19#go-unix
package main

import "fmt"

func main() {
	fmt.Println("run on unix")
}

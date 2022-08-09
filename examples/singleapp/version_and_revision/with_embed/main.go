/*
Go 1.16 から追加された embed パッケージを利用して内部の変数に外部ファイルデータを埋め込むサンプルです.

Commands:

	$ git describe --tags --abbrev=0 > version.txt
	$ git rev-list -1 HEAD > revision.txt
	$ git describe --tags > build.txt
	$ go run .
*/
package main

import (
	_ "embed"
	"fmt"
	"os"
)

// Version and Revision
//
// Commands:
//
//	$ git describe --tags --abbrev=0 > version.txt
//	$ git rev-list -1 HEAD > revision.txt
//	$ git describe --tags > build.txt
var (
	//go:embed version.txt
	version string
	//go:embed revision.txt
	revision string
	//go:embed build.txt
	build string
)

func main() {
	os.Exit(run())
}

func run() int {
	fmt.Printf("Version : %s", version)
	fmt.Printf("Revision: %s", revision)
	fmt.Printf("Build   : %s", build)

	return 0
}

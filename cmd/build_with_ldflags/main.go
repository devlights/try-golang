/*
	go build 時に -ldflags を指定して内部の変数に外部から値を注入するサンプルです.

	例:
		$ go build -race -ldflags "-X main.version=$(git describe --tag --abbrev=0) -X main.revision=$(git rev-list -1 HEAD)"
		$ go run  -race -ldflags "-X main.version=$(git describe --tag --abbrev=0) -X main.revision=$(git rev-list -1 HEAD)" .
*/
package main

import (
	"fmt"
	"os"
)

// Version and Revision
//
// REFERENCES:
//   - https://blog.alexellis.io/inject-build-time-vars-golang/
//   - https://christina04.hatenablog.com/entry/2016/12/08/101114
//   - https://stackoverflow.com/questions/1404796/how-to-get-the-latest-tag-name-in-current-branch-in-git
//     - https://stackoverflow.com/a/7261049
var (
	version  string
	revision string
)

func main() {
	os.Exit(run())
}

func run() int {
	fmt.Printf("Version : %s\n", version)
	fmt.Printf("Revision: %s\n", revision)
	return 0
}

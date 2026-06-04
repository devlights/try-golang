/*
go build 時に -ldflags を指定して内部の変数に外部から値を注入するサンプルです.

例:

	$ go build -race -ldflags \
		" \
			-X main.version=$(git describe --tag --abbrev=0) \
			-X main.revision=$(git rev-list -1 HEAD) \
			-X main.build=$(git describe --tags) \
		"
	$ go run  -race -ldflags \
		" \
			-X main.version=$(git describe --tag --abbrev=0) \
			-X main.revision=$(git rev-list -1 HEAD) \
			-X main.build=$(git describe --tags) \
		" .
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
//   - https://stackoverflow.com/a/7261049
//   - https://git-scm.com/book/ja/v2/Appendix-C%3A-Git%E3%81%AE%E3%82%B3%E3%83%9E%E3%83%B3%E3%83%89-%E6%A4%9C%E6%9F%BB%E3%81%A8%E6%AF%94%E8%BC%83
//   - https://git-scm.com/book/ja/v2/Git-%E3%81%A7%E3%81%AE%E5%88%86%E6%95%A3%E4%BD%9C%E6%A5%AD-%E3%83%97%E3%83%AD%E3%82%B8%E3%82%A7%E3%82%AF%E3%83%88%E3%81%AE%E9%81%8B%E5%96%B6#r_build_number
var (
	version  string
	revision string
	build    string
)

func main() {
	os.Exit(run())
}

func run() int {
	fmt.Printf("Version : %s\n", version)
	fmt.Printf("Revision: %s\n", revision)
	fmt.Printf("Build   : %s\n", build)
	return 0
}

/*
	Go 1.16 から追加された embed パッケージを利用して内部の変数に外部ファイルデータを埋め込むサンプルです.

	Commands:
		$ git describe --tags --abbrev=0 > version.txt
		$ git rev-list -1 HEAD >> version.txt
		$ git describe --tags >> version.txt
		$ go run -race .
*/
package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"os"
)

type version []byte

func (v version) String() string {
	buf := bytes.NewBuffer(v)
	get := func() string { v, _ := buf.ReadString('\n'); return v }

	return fmt.Sprintf("Version : %sRevision: %sBuild   : %s", get(), get(), get())
}

//go:embed version.txt
var v []byte

func main() {
	os.Exit(run())
}

func run() int {
	fmt.Print(version(v))
	return 0
}

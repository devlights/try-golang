/*
Go の標準パッケージだけを使って HTTP サーバをローカルで立てるサンプル

REFERENCES:
  - https://pkg.go.dev/net/http
  - https://hodalog.com/generate-self-signed-certificate-using-by-golang/
  - https://fm-cowkey.hatenablog.com/entry/2018/01/27/154721
  - https://code-database.com/knowledges/87
  - https://zenn.dev/tomi/articles/2020-10-02-go-web3
*/
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	ret := 0

	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		ret = 1
	}

	os.Exit(ret)
}

func run() error {
	http.Handle("/", http.FileServer(http.Dir("html/")))
	log.Println(http.ListenAndServe(":8888", nil))

	return nil
}

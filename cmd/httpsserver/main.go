/*
Go の標準パッケージだけを使って HTTPS サーバをローカルで立てるサンプル

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
	//
	// Go の標準パッケージの中にオレオレ証明書を作成するためのソースがある
	// 以下の様にして呼び出せる
	//
	// $ $(go env GOROOT)/src/crypto/tls/generate_cert.go -rsa-bits 2048 -host localhost
	//
	// 実行すると、key.pem と cert.pem ファイルが生成される
	//
	// 後は、このファイルを使ってサーバーを起動するだけとなる。
	// ただし、正式な証明書ではないため、そのままChromeで開くと証明書エラーが表示されるので
	//   chrome://flags/#allow-insecure-localhost
	// をブラウザのURL欄に入力し
	//   Allow invalid certificates for resources loaded from localhost
	// の項目を Enabled に設定する。
	//
	http.Handle("/", http.FileServer(http.Dir("html/")))
	log.Println(http.ListenAndServeTLS(":8888", "./cert.pem", "./key.pem", nil))

	return nil
}

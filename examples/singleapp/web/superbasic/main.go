/*
Go webアプリ サンプル （超基本）

REFERENCES:
  - https://pkg.go.dev/net/http@latest
  - https://astaxie.gitbooks.io/build-web-application-with-golang/content/ja/
*/
package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from: %s", r.Header.Get("Referer"))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9999", nil)
}

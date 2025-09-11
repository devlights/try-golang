package main

import (
	_ "embed"
	"log"
	"net/http"

	"github.com/devlights/try-golang/examples/http/01.basic/handler/index"
)

var (
	//go:embed index/view.html
	html []byte
)

func main() {
	http.Handle("/", index.Handler(html))

	if err := http.ListenAndServe(":8888", nil); err != nil {
		log.Fatal(err)
	}
}

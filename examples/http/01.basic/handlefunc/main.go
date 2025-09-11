package main

import (
	_ "embed"
	"log"
	"net/http"
)

var (
	//go:embed index.html
	html []byte
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write(html); err != nil {
			log.Fatal(err)
		}
	})

	if err := http.ListenAndServe(":8888", nil); err != nil {
		log.Fatal(err)
	}
}

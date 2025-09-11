package main

import (
	"fmt"
	"net/http"
	"time"
)

// # REFERENCE
//   - https://zenn.dev/tady/articles/28fd15b62f3767
func main() {
	var (
		ch = make(chan int64)
	)
	defer close(ch)

	http.HandleFunc("/polling", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("%d\n", <-ch)))
	})

	http.HandleFunc("/send", func(w http.ResponseWriter, r *http.Request) {
		ch <- time.Now().Unix()
		w.Write([]byte("OK\n"))
	})

	http.ListenAndServe(":8888", nil)
}

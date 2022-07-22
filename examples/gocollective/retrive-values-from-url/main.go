// Stackoverflow Go Collective example
//
// How to retrieve values from URL
//
// URL
//   - https://stackoverflow.com/questions/73079531/how-to-retrieve-values-from-the-url-in-go
//
// REFERENCES
//   - https://pkg.go.dev/net/http@latest
//   - https://stackoverflow.com/questions/39320025/how-to-stop-http-listenandserve
package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	var (
		mux      = http.NewServeMux()
		srv      = &http.Server{Addr: ":8888", Handler: mux}
		ctx, cxl = context.WithTimeout(context.Background(), 3*time.Second)
	)
	defer cxl()

	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		fmt.Println(query.Get("message"))
	})

	go func() {
		srv.ListenAndServe()
	}()

	go func() {
		for {
			_, err := http.Get("http://localhost:8888/hello?message=world")
			if err == nil {
				break
			}
		}
	}()

	<-ctx.Done()
	srv.Shutdown(ctx)
}

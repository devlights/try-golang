package main

import (
	"context"
	"log"
	"net/http"
	"time"
)

func main() {
	var (
		mux              = http.NewServeMux()
		srv              = &http.Server{Addr: ":8888", Handler: mux}
		mainCtx          = context.Background()
		procCtx, procCxl = context.WithTimeout(mainCtx, 3*time.Second)
	)
	defer procCxl()

	// ----------- Start ----------- //

	go func() {
		log.Println("Server is up.")
		srv.ListenAndServe()
	}()

	<-procCtx.Done()

	// ----------- Shutdown ----------- //

	var (
		shutdownCtx, shutdownCxl = context.WithTimeout(mainCtx, 1*time.Second)
	)
	defer shutdownCxl()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		switch err {
		case context.DeadlineExceeded:
			log.Println("Server shutdown process timed out.")
		default:
			log.Fatal(err)
		}
	}
	log.Println("Server has been shutdown.")
}

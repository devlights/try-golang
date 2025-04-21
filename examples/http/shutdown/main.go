package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	log.SetFlags(log.Lmicroseconds)

	if err := run(); err != nil {
		log.Fatalf("Application failed: %v", err)
	}
}

func run() error {
	var (
		port string
	)
	if port = os.Getenv("PORT"); port == "" {
		port = "8888"
	}

	var (
		mux = http.NewServeMux()
		srv = &http.Server{
			Addr:           ":" + port,
			Handler:        mux,
			MaxHeaderBytes: 1 << 20, // 1MB
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			IdleTimeout:    120 * time.Second,
		}
		mainCtx, mainCxl = context.WithCancel(context.Background())
		errs             = make(chan error, 1)
	)
	defer mainCxl()

	// ----------- Handlers ----------- //
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World")
	})

	// ----------- Start ----------- //
	go func() {
		log.Printf("Server is listening on %s", srv.Addr)
		if err := srv.ListenAndServe(); err != nil {
			errs <- err
		}
	}()

	// ----------- Signals ----------- //
	var (
		sigs = make(chan os.Signal, 1)
	)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer signal.Stop(sigs)

	// ----------- Selects ----------- //
	select {
	case err := <-errs:
		if !errors.Is(err, http.ErrServerClosed) {
			return fmt.Errorf("server error: %w", err)
		}
		return nil
	case sig := <-sigs:
		log.Printf("Shutdown started: %v", sig)

		// ----------- Shutdown ----------- //
		shutCtx, shutCxl := context.WithTimeout(mainCtx, 5*time.Second)
		defer shutCxl()

		if err := srv.Shutdown(shutCtx); err != nil {
			if errors.Is(err, context.DeadlineExceeded) {
				log.Println("Server shutdown process timed out.")
			} else {
				log.Printf("Server shutdown failed: %v", err)
			}

			if closeErr := srv.Close(); closeErr != nil {
				return fmt.Errorf("force close failed: %w", closeErr)
			}

			return fmt.Errorf("shutdown error: %w", err)
		}

		log.Println("Server has been gracefully shutdown.")

		return nil
	}
}

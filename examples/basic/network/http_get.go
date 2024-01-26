package network

import (
	"bytes"
	"errors"
	"io"
	"net/http"

	"github.com/devlights/gomy/output"
)

// HttpGet -- http.Getを使ったサンプルです.
func HttpGet() error {
	var (
		httpCh = make(chan []byte)
		dataCh = make(chan []byte)
		errCh  = make(chan error, 1)
	)
	defer close(errCh)

	// [goroutine-1] Call http.Get and send response body to http-channel
	go func() {
		defer close(httpCh)

		const (
			url = "https://github.com/devlights/try-golang"
		)

		resp, err := http.Get(url)
		if err != nil {
			errCh <- err
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			errCh <- err
			return
		}
		httpCh <- body

		output.Stderrl("recv bytes", len(body))
	}()

	// [goroutine-2] Recv from http-channel and process http response
	go func() {
		defer close(dataCh)

		var (
			sTag = []byte("<title>")
			eTag = []byte("</title>")
		)

		for body := range httpCh {
			var (
				start = bytes.Index(body, sTag)
				end   = bytes.Index(body, eTag)
			)

			if start == -1 || end == -1 {
				errCh <- errors.New("title tag does not exist")
				return
			}

			var (
				buf = body[start+len(sTag) : end]
			)
			dataCh <- buf
		}
	}()

	// output
	select {
	case err := <-errCh:
		return err
	case b := <-dataCh:
		output.Stderrf("http.get", "%s\n", b)
	}

	return nil

	/*
	   $ task
	   task: Task "build" is up to date
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: network_http_get

	   [Name] "network_http_get"
	   recv bytes           173143
	   http.get             GitHub - devlights/try-golang: This is my TUTORIAL project for golang.


	   [Elapsed] 448.931025ms
	*/

}

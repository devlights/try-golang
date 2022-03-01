package network

import (
	"bytes"
	"io"
	"net/http"
	"sync"

	"github.com/devlights/gomy/output"
)

// HttpGet -- http.Getを使ったサンプルです.
func HttpGet() error {
	var (
		wg     sync.WaitGroup
		httpCh = make(chan []byte, 1)
		dataCh = make(chan []byte, 1)
		errCh  = make(chan error, 1)
	)

	// [goroutine-1] Call http.Get and send bytes to http-channel
	wg.Add(1)
	go func() {
		defer wg.Done()
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

		b, err := io.ReadAll(resp.Body)
		if err != nil {
			errCh <- err
			return
		}
		httpCh <- b

		output.Stdoutl("recv bytes", len(b))
	}()

	// [goroutine-2] Recv from http-channel and filtering http response
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(dataCh)

		for b := range httpCh {
			var (
				start = bytes.Index(b, []byte("<title>"))
				end   = bytes.Index(b, []byte("</title>"))
				buf   = b[start+len("<title>") : end]
			)
			dataCh <- buf
		}
	}()

	// [goroutine-3] Close the channel when the task is complete
	go func() {
		defer close(errCh)
		wg.Wait()
	}()

	// has error?
	for e := range errCh {
		return e
	}

	// output
	for b := range dataCh {
		output.Stdoutf("http.get -- title", "%s\n", b)
	}

	return nil
}

package readwrite

import (
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

// NullWriter -- io.Discard のサンプルです.
//
// REFERENCES:
//   - https://golang.org/io/#pkg-variables
//   - https://stackoverflow.com/a/25344458
func NullWriter() error {
	// ----------------------------------------------------------------
	// io.Discard は、io.Writer を実装しているけど何もしません。
	// 処理を行う上で io.Writer が必要だが、その結果は必要ない場合などに利用します。
	// ----------------------------------------------------------------
	var (
		urls = []string{
			"https://www.alexa.com/topsites/", // このサイトは大きいので取得に時間がかかる
			"https://www.google.co.jp/",
			"https://golang.org/",
			"https://github.com/",
		}
	)

	var (
		logger = log.New(os.Stdout, "", 0)
		wg     sync.WaitGroup
		errCh  = make(chan error, len(urls))
	)

	var (
		fetch = func(wg *sync.WaitGroup, url string) {
			defer wg.Done()
			defer func(start time.Time) {
				logger.Printf("fetch: %-40s --> %v\n", url, time.Since(start))
			}(time.Now())

			var (
				client = http.Client{
					Timeout: 1 * time.Second,
				}
			)

			resp, err := client.Get(url)
			if err != nil {
				errCh <- err
				return
			}

			defer func() {
				if err = resp.Body.Close(); err != nil {
					errCh <- err
				}
			}()

			// 結果は必要ないので捨てる
			_, err = io.Copy(io.Discard, resp.Body)
			if err != nil {
				errCh <- err
				return
			}
		}
	)

	start := time.Now()
	for _, url := range urls {
		wg.Add(1)
		go fetch(&wg, url)
	}

	wg.Wait()
	close(errCh)

	for e := range errCh {
		logger.Printf("Error: %v\n", e)
	}

	logger.Printf("elapsed: %v\n", time.Since(start))

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: fileio_null_writer

	   [Name] "fileio_null_writer"
	   fetch: https://www.google.co.jp/                --> 73.324262ms
	   fetch: https://github.com/                      --> 75.424952ms
	   fetch: https://golang.org/                      --> 300.619817ms
	   fetch: https://www.alexa.com/topsites/          --> 325.817194ms
	   elapsed: 325.983864ms


	   [Elapsed] 325.999604ms
	*/

}

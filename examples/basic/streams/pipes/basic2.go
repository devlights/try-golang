package pipes

import (
	"bufio"
	"io"
	"sync"

	"github.com/devlights/gomy/output"
)

// Basic2 -- io.Pipe() のサンプルです。
//
// #REFERENCES
//   - https://www.geeksforgeeks.org/io-pipe-function-in-golang-with-examples/
//   - https://medium.com/eureka-engineering/file-uploads-in-go-with-io-pipe-75519dfa647b
func Basic2() error {
	var (
		wg   sync.WaitGroup
		errs = make(chan error, 2)
	)

	// ------------------------------------------------
	// Pipeの読み取り側と書き込み側を取得
	// ------------------------------------------------

	var (
		pr, pw = io.Pipe()
	)

	// ------------------------------------------------
	// 書き込み側
	//  ------------------------------------------------

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer pw.Close()

		for _, v := range []string{"hello", "world"} {
			_, err := pw.Write([]byte(v))
			if err != nil {
				errs <- err
				return
			}
		}
	}()

	// ------------------------------------------------
	// 読み込み側
	//  ------------------------------------------------

	wg.Add(1)
	go func() {
		defer wg.Done()

		scanner := bufio.NewScanner(pr)
		for scanner.Scan() {
			output.Stdoutl("[pr]", scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			errs <- err
		}
	}()

	// ------------------------------------------------
	// 非同期タスクが完了するまで待機
	// ------------------------------------------------

	wg.Wait()

	close(errs)
	for e := range errs {
		return e
	}

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: streams_pipe_basic2

	   [Name] "streams_pipe_basic2"
	   [pr]                 helloworld


	   [Elapsed] 252.58µs
	*/

}

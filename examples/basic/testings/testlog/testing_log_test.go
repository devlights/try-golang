package testlog

import (
	"fmt"
	"testing"
	"time"
)

func TestTLog(t *testing.T) {
	// ----------------------------------------------------------------------
	// t.Log() の挙動が go1.13 と go1.14 では以下のように変わっている
	//
	// [go1.13]
	//   - バッファリングされていて、テストが完了した後に出力される
	// [go1.14]
	//   - バッファリング無しで、即出力される
	//
	// REFERENCES::
	//   - https://dave.cheney.net/2020/03/10/go-test-v-streaming-output
	//   - https://devlights.hatenablog.com/entry/2019/11/07/003735
	// ----------------------------------------------------------------------
	for i := 0; i < 5; i++ {
		fmt.Printf("fmt %d\n", i)
		t.Logf("t.Log %d\n", i)

		time.Sleep(100 * time.Millisecond)
	}

	/*
		$ go version
		go version go1.14 windows/amd64

		$ cd $(go env GOPATH)
		$ go get golang.org/dl/go1.13
		$ go1.13 download
		Downloaded   0.0% (    16384 / 133885977 bytes) ...
		Downloaded   0.7% (   899869 / 133885977 bytes) ...
		Downloaded   1.2% (  1669917 / 133885977 bytes) ...
		・
		・
		・
		Downloaded 100.0% (133885977 / 133885977 bytes)
		Success. You may now run 'go1.13'

		$ go1.13 version
		go version go1.13 windows/amd64

		$ go1.13 test -v examples/basic/testings/testing_log_test.go
		=== RUN   TestTLog
		fmt 0
		fmt 1
		fmt 2
		fmt 3
		fmt 4
		--- PASS: TestTLog (5.00s)
		    testing_log_test.go:20: t.Log 0
		    testing_log_test.go:20: t.Log 1
		    testing_log_test.go:20: t.Log 2
		    testing_log_test.go:20: t.Log 3
		    testing_log_test.go:20: t.Log 4
		PASS
		ok      command-line-arguments  5.095s


		$ go test -v examples/basic/testings/testing_log_test.go
		=== RUN   TestTLog
		fmt 0
		    TestTLog: testing_log_test.go:20: t.Log 0
		fmt 1
		    TestTLog: testing_log_test.go:20: t.Log 1
		fmt 2
		    TestTLog: testing_log_test.go:20: t.Log 2
		fmt 3
		    TestTLog: testing_log_test.go:20: t.Log 3
		fmt 4
		    TestTLog: testing_log_test.go:20: t.Log 4
		--- PASS: TestTLog (5.00s)
		PASS
		ok      command-line-arguments  5.086s
	*/
}

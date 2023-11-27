package rand

import (
	"crypto/rand"

	"github.com/devlights/gomy/iter"
	"github.com/devlights/gomy/output"
)

// Read -- crypto/rand.Read を用いてセキュリティ的に安全な乱数を生成するサンプルです.
//
// # REFERENCES
//   - https://pkg.go.dev/crypto/rand@go1.19#pkg-variables
//   - https://github.com/golang/go/wiki/CodeReviewComments#crypto-rand
func Read() error {

	for range iter.Range(10) {
		var (
			buf = make([]byte, 32)
			err error
		)

		if _, err = rand.Read(buf); err != nil {
			return err
		}

		output.Stdoutf("[Crypto][rand.Read]", "%x\n", buf)
	}

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: crypto_rand_read

	   [Name] "crypto_rand_read"
	   [Crypto][rand.Read]  5e1d604a62647e940ffd655fde24783e2d95de503d31e4272011a02b5456d544
	   [Crypto][rand.Read]  8711bba87e9374c02e5e6aa128a64bc2b314bd5da6d0340d05e1cb70397cb4b6
	   [Crypto][rand.Read]  39a744bbb8b46a6a6babba3223ff10a88981a01157317cda57e748d5e7070a22
	   [Crypto][rand.Read]  bf5777bc2b1692b0eaca8a17b7776f3ce79631a93ac4d6f46ac3d1030268e209
	   [Crypto][rand.Read]  c39ce5f19ca6b801acd81c849944457f2adfd8e981b89e726cb51270ac51ac46
	   [Crypto][rand.Read]  1d65a859a3f8bde8d5c5ef123ce31d92b17a0a055be254757830e015a374ea83
	   [Crypto][rand.Read]  e9c0fb22876ca5f1cda72251d574e07dc4056751bbb7e8865818fd930bd31ada
	   [Crypto][rand.Read]  9434bfb5f579398787ad78f0a5d3c6ca752331b61285bdb015451c9203b44802
	   [Crypto][rand.Read]  216f6f34ed7468d6f7ba97ba63523578336eac29372587aeb840579ed50e9841
	   [Crypto][rand.Read]  30cc31ca66a26c97585400bfa1e654a21373125c9936b91bca8b943746f375fe


	   [Elapsed] 118.88µs
	*/

}

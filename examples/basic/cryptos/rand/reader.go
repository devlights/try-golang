package rand

import (
	"crypto/rand"
	"encoding/hex"
	"io"

	"github.com/devlights/gomy/iter"
	"github.com/devlights/gomy/output"
)

// Reader -- crypto/rand.Reader を用いてセキュリティ的に安全な乱数を生成するサンプルです.
//
// # REFERENCES
//   - https://pkg.go.dev/crypto/rand@go1.19#pkg-variables
//   - https://github.com/golang/go/wiki/CodeReviewComments#crypto-rand
func Reader() error {
	var (
		r = rand.Reader
	)

	for range iter.Range(10) {
		var (
			b   = make([]byte, 32)
			err error
		)

		_, err = io.ReadFull(r, b)
		if err != nil {
			return err
		}

		output.Stdoutl("[Crypto][rand.Reader]", hex.EncodeToString(b))
	}

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: crypto_rand_reader

	   [Name] "crypto_rand_reader"
	   [Crypto][rand.Reader] 307f9d04d58969c2678e0e3a88b4061778f4dc8857364f6eb17feb62336db65e
	   [Crypto][rand.Reader] 23a020127adb882746c8d50b9fd71921be27c9f7788eb5bd961a379703c4fe6e
	   [Crypto][rand.Reader] 45739916586e01eb5680417ebbb2484bdf002e12f963e30fd36bb15f93de8892
	   [Crypto][rand.Reader] 92d2f9b8ae758f488c8fc0fdb1a41b33f9c2fda8afa2f81eb24c6942929dc8d9
	   [Crypto][rand.Reader] a6202c11c9607c44b86f3152f10305369002a407e779dca8013a4de5e47c726c
	   [Crypto][rand.Reader] ec4b91bf90c6b8e26670820e1f047d831c71606f72f5af1122bc983969b07291
	   [Crypto][rand.Reader] ed830f6ccf4bfa2f4f279a2fa8c55baac7eb4f34a8b2314cb84e9c8b06b73f00
	   [Crypto][rand.Reader] 987b43ade967123324776d334d51f17e736f290ade08772dc1db00527c233e9f
	   [Crypto][rand.Reader] c273e8ded06028d4459ebc71553b4ac002868934b6ec1bca764b31a8f7426c0a
	   [Crypto][rand.Reader] 1e716242912aa9ac145c96e554c70c50a3a6d1687813956b967c44bbde4b4de0


	   [Elapsed] 107.62µs
	*/

}

package strs

import (
	"math/rand"
	"time"

	"github.com/devlights/gomy/output"
)

// RandomString は、指定された文字数のランダム文字列を作成するサンプルです.
//
// # REFERENCES
//   - https://gist.github.com/devlights/7534500bfe62c566bf944553ae8974e8
func RandomString() error {
	const (
		charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	)

	var (
		buf       = make([]byte, 1<<6)
		unixNano  = time.Now().UnixNano()
		rndSource = rand.NewSource(unixNano)
		rnd       = rand.New(rndSource)
	)

	for i := range buf {
		buf[i] = charset[rnd.Intn(len(charset))]
	}

	output.Stdoutl("[output]", string(buf))

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: string_random_string

	   [Name] "string_random_string"
	   [output]             0OWayZgY57QSJKHvAl8ePRtFSFGtgKJRug6OFdN3XL17oxUW2pqmCaGpGqYV2oEY

	   [Elapsed] 28.55µs
	*/
}

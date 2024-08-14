package timeoutreader

import (
	"math/rand"
	"time"
)

const (
	charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func randomString(buf []byte) {
	var (
		unixNano  = time.Now().UnixNano()
		rndSource = rand.NewSource(unixNano)
		rnd       = rand.New(rndSource)
	)

	for i := range buf {
		buf[i] = charset[rnd.Intn(len(charset))]
	}
}

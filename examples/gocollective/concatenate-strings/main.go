// Stackoverflow Go Collective example
//
// # How to efficiently concatenate strings in go
//
// URL
//   - https://stackoverflow.com/questions/1760757/how-to-efficiently-concatenate-strings-in-go
//
// REFERENCES
//   - https://pkg.go.dev/strings@latest#Builder
//   - https://yourbasic.org/golang/measure-execution-time/
package main

import (
	"bytes"
	"log"
	"strings"
	"time"

	"github.com/devlights/gomy/times"
)

func init() {
	log.SetFlags(0)
}

func main() {
	// Go 1.10 より strings.Builder が追加された。
	// bytes.Buffer でも良いが、文字列の場合はこちらのほうが効率が良い。

	var (
		sb  strings.Builder
		buf bytes.Buffer
	)

	// ----- bytes.Buffer ----- //
	elapsed := times.Stopwatch(func(start time.Time) {
		for i := 0; i < 3_000_000; i++ {
			buf.WriteString("i")
		}
	})
	log.Printf("[bytes.Buffer   ] len=%d, elapsed=%v", buf.Len(), elapsed)

	// ----- strings.Builder ----- //
	elapsed = times.Stopwatch(func(start time.Time) {
		for i := 0; i < 3_000_000; i++ {
			sb.WriteString("i")
		}
	})
	log.Printf("[strings.Builder] len=%d, elapsed=%v", sb.Len(), elapsed)
}

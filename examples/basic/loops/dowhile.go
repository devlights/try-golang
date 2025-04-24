package loops

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

// DoWhile は、Goで他の言語にある do-while と同様のことを行うサンプルです.
//
// Goにはループ制御が for しかないので、他の言語にある do-while は当然構文としては用意されていない。
// しかし、for で同じことは当然出来る。
func DoWhile() error {
	type (
		ctxKey struct{}
	)
	const (
		retryMax = 2
	)
	var (
		l   = log.New(os.Stdout, "", log.Lmicroseconds)
		i   = 0
		ctx = context.Background()

		fn = func(ctx context.Context) {
			// なんか処理している風
			_, err := net.DialTimeout("tcp", ":12345", 10*time.Millisecond)
			l.Printf("[%d] %v", ctx.Value(ctxKey{}).(int), err)
		}
		hr = func() {
			fmt.Println("--------------------------------")
		}
	)

	// do-while の代替 (1)
	for {
		fn(context.WithValue(ctx, ctxKey{}, i))
		i++

		// do-whileの条件判定の代わり
		if i > retryMax {
			break
		}
	}

	hr()

	// do-while の代替 (2)
	i = 0
	for ok := true; ok; ok = (i <= retryMax) {
		fn(context.WithValue(ctx, ctxKey{}, i))
		i++
	}

	return nil
}

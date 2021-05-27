package goroutines

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/devlights/gomy/chans"
)

// WithContextDeadline -- context.WithDeadline を使ったサンプルです
func WithContextDeadline() error {
	// ------------------------------------------------------------------
	// context.WithDeadline() は、指定した time.Time に到達した時点で
	// 完了となるコンテキストを返す. context.WithTimeout() と似た動き.
	// ------------------------------------------------------------------
	const (
		goroutineCount = 2
	)

	// utility functions
	var (
		iter     = func(n int) []struct{} { return make([]struct{}, n) }
		toDoneCh = func(ctxs ...context.Context) []<-chan struct{} {
			dones := make([]<-chan struct{}, 0, len(ctxs))

			for _, c := range ctxs {
				dones = append(dones, c.Done())
			}

			return dones
		}
	)

	// logger
	var (
		mainLog = log.New(os.Stdout, "[main] ", 0)
		gLog    = log.New(os.Stderr, "[goroutine] ", 0)
	)

	// deadline
	var (
		now      = time.Now()
		deadline = now.Add(3 * time.Second)
	)

	// contexts
	var (
		rootCtx             = context.Background()
		mainCtx, mainCancel = context.WithCancel(rootCtx)
		procCtx, procCancel = context.WithTimeout(mainCtx, 5*time.Second)
	)

	defer mainCancel()
	defer procCancel()

	timelimit, _ := procCtx.Deadline()
	mainLog.Printf("start now=%d\tdeadline=%d\ttimelimit=%d", now.UTC().Unix(), deadline.UTC().Unix(), timelimit.UTC().Unix())

	// start goroutines
	ctxs := make([]context.Context, 0, goroutineCount)
	for i := range iter(goroutineCount) {
		ctxs = append(ctxs, func(no int) context.Context {
			ctx, cancel := context.WithDeadline(procCtx, deadline)

			go func() {
				defer cancel()

			LOOP:
				for {
					select {
					case <-ctx.Done():
						break LOOP
					default:
						gLog.Println(no, time.Now().UTC().Unix())
					}

					select {
					case <-ctx.Done():
						break LOOP
					case <-time.After(1 * time.Second):
					}
				}
			}()

			return ctx
		}(i))
	}

	// wait until all goroutines is done.
	<-chans.WhenAll(toDoneCh(ctxs...)...)

	mainLog.Println("all goroutines done", time.Now().UTC().Unix())
	mainLog.Println("done")

	return nil
}

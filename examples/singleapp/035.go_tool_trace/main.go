// go tool trace のサンプルです.
//
// # REFERENCES
//   - https://youngstone89.medium.com/go-tools-trace-25e4c1c442ff
//   - https://budougumi0617.github.io/2020/12/04/goroutine_tuning_with_benchmark_benchstat_trace/
package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime/trace"
	"time"

	"github.com/devlights/gomy/chans"
	"github.com/devlights/gomy/ctxs"
	"github.com/devlights/gomy/errs"
)

func main() {
	trace.Start(os.Stderr)
	defer trace.Stop()

	var (
		rootCtx  = context.Background()
		ctx, cxl = context.WithTimeout(rootCtx, 10*time.Second)
	)
	defer cxl()

	var (
		urls = []string{
			"https://devlights.hatenablog.com/",
			"https://github.com/devlights/try-golang",
			"https://github.com/devlights/gomy",
			"https://github.com/devlights/goxcel",
			"https://github.com/devlights/try-csharp",
			"https://github.com/devlights/try-python",
		}
		in     = chans.GeneratorContext(ctx, urls...)
		wCnt   = 3
		fanout = chans.FanOutContext(ctx, in, wCnt, func(url string) {
			var (
				resp = errs.Stderr(http.Get(url))
				buf  = errs.Stderr(io.ReadAll(resp.Body))
			)
			defer resp.Body.Close()

			fmt.Printf("%s [%dbytes]\n", url, len(buf))
		})
	)

	<-ctxs.WhenAll(ctx, fanout...).Done()
}

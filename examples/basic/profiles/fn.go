package profiles

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/devlights/gomy/chans"
	"github.com/devlights/gomy/ctxs"
	"github.com/devlights/gomy/errs"
)

func run(pCtx context.Context) context.Context {
	var (
		ctx, cxl = context.WithCancel(pCtx)
	)

	var (
		urls = []string{
			"https://devlights.hatenablog.com/",
			"https://github.com/devlights/try-golang",
			"https://github.com/devlights/gomy",
			"https://github.com/devlights/goxcel",
			"https://github.com/devlights/try-csharp",
			"https://github.com/devlights/try-python",
		}
		in          = chans.GeneratorContext(ctx, urls...)
		workerCount = len(urls)
		fanout      = chans.FanOutContext(ctx, in, workerCount, func(url string) {
			var (
				resp = errs.Stderr(http.Get(url))
				buf  = errs.Stderr(io.ReadAll(resp.Body))
			)
			defer resp.Body.Close()
			fmt.Printf("%s [%dbytes]\n", url, len(buf))
		})
	)

	go func() {
		defer cxl()
		<-ctxs.WhenAll(ctx, fanout...).Done()
	}()

	return ctx
}

package main

import (
	"flag"
	"log/slog"
	"os"
)

func main() {
	discard := flag.Bool("discard", false, "Use slog.DiscardHandler")
	flag.Parse()

	if err := run(*discard); err != nil {
		panic(err)
	}
}

func run(discard bool) error {
	//
	// Go 1.24 にて、slog.DiscardHandler が追加された。
	// io.Discard と同様に、このハンドラは出力を行わない。
	//
	// > DiscardHandler discards all log output. DiscardHandler.Enabled returns false for all Levels.
	// > (DiscardHandler は、すべてのログ出力を破棄します。DiscardHandler.Enabled は、すべての Levels に対して false を返します。)
	//
	// # REFERENCES
	//   - https://pkg.go.dev/log/slog#example-package-DiscardHandler
	//
	var (
		l = slog.New(handler(discard))
	)

	l.Info("helloworld")
	l.Warn("helloworld")
	l.Error("helloworld")

	return nil
}

func handler(discard bool) slog.Handler {
	if discard {
		return slog.DiscardHandler
	}

	return slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{ReplaceAttr: rmTime})
}

// From https://cs.opensource.google/go/go/+/refs/tags/go1.24.0:src/log/slog/internal/slogtest/slogtest.go;l=13
func rmTime(groups []string, a slog.Attr) slog.Attr {
	if a.Key == slog.TimeKey && len(groups) == 0 {
		return slog.Attr{}
	}
	return a
}

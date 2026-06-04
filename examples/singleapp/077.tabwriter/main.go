package main

import (
	"bufio"
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"slices"
	"text/tabwriter"
	"time"
)

var (
	rootCtx = context.Background()
	errCh   = make(chan error)
)

func main() {
	log.SetFlags(0)

	var (
		ctx, cxl = context.WithTimeout(rootCtx, 1*time.Second)
		err      error
	)
	defer cxl()

	if err = run(ctx); err != nil {
		log.Panic(err)
	}
}

func run(pCtx context.Context) error {
	var (
		ctx, cxl = context.WithTimeout(pCtx, 300*time.Millisecond)
		err      error
	)
	defer cxl()

	var (
		items   = gen(ctx)               // 処理対象を生成し
		tsv     = toTsv(ctx, items)      // tsvに変換して (encoding/csv   を利用)
		aligned = toTabAligned(ctx, tsv) // タブで揃える  (text/tabwriter を利用)
	)
	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("run: %w", ctx.Err())
		case err = <-errCh:
			return err
		case v, ok := <-aligned:
			if !ok {
				return nil
			}

			log.Println(v)
		}
	}
}

// gen は、処理対象を生成します。
func gen(pCtx context.Context) <-chan []string {
	var (
		out = make(chan []string)
	)
	go func() {
		defer close(out)

		var (
			items = [][]string{
				{"a", "b", "c"},
				{"aa", "bb", "cc"},
				{"aaa", "bbb", "ccc"},
				{"123456", "1234567", "12345678"},
			}
		)
		for item := range slices.Values(items) {
			select {
			case <-pCtx.Done():
				errCh <- fmt.Errorf("gen: %w", pCtx.Err())
				return
			default:
				out <- item
			}
		}
	}()

	return out
}

// toTsv は、データをTSV化します。
func toTsv(pCtx context.Context, in <-chan []string) <-chan string {
	var (
		out       = make(chan string)
		pr, pw, _ = os.Pipe()
	)
	go func() {
		var (
			csv = csv.NewWriter(pw)
		)
		defer pw.Close()
		defer csv.Flush()

		csv.Comma = '\t'

		for item := range in {
			select {
			case <-pCtx.Done():
				return
			default:
				if err := csv.Write(item); err != nil {
					errCh <- fmt.Errorf("toTsv: %w", err)
					return
				}
			}
		}
	}()
	go func() {
		defer close(out)

		var (
			scanner = bufio.NewScanner(pr)
		)
		for scanner.Scan() {
			select {
			case <-pCtx.Done():
				return
			default:
				out <- fmt.Sprintf("%s\t", scanner.Text()) // 末尾にタブが無いとtabwriterが調整してくれない
			}

			if err := scanner.Err(); err != nil {
				errCh <- fmt.Errorf("toTsv: %w", err)
				return
			}
		}
	}()

	return out
}

// toTabAligned は、タブで揃えた出力に整形します。
//
// text/tabwriterパッケージを用いて整形処理を行っています。
// tabwriterはCLIアプリで出力する際に便利。データとして扱う場合は encoding/csv を利用します。
func toTabAligned(pCtx context.Context, in <-chan string) <-chan string {
	var (
		out       = make(chan string)
		pr, pw, _ = os.Pipe()
	)
	go func() {
		var (
			minwidth = 8
			tabwidth = 4
			padding  = 0
			padchar  = '.'
			flags    = tabwriter.AlignRight | tabwriter.Debug
			tw       = tabwriter.NewWriter(pw, minwidth, tabwidth, padding, byte(padchar), flags)
		)
		defer pw.Close()
		defer tw.Flush()

		for v := range in {
			select {
			case <-pCtx.Done():
				errCh <- fmt.Errorf("toTabAligned: %w", pCtx.Err())
				return
			default:
				if _, err := tw.Write(fmt.Appendln(nil, v)); err != nil {
					errCh <- fmt.Errorf("toTabAligned: %w", err)
					return
				}
			}
		}
	}()
	go func() {
		defer close(out)

		var (
			scanner = bufio.NewScanner(pr)
		)
		for scanner.Scan() {
			select {
			case <-pCtx.Done():
				errCh <- fmt.Errorf("toTabAligned: %w", pCtx.Err())
				return
			default:
				out <- scanner.Text()
			}

			if err := scanner.Err(); err != nil {
				errCh <- fmt.Errorf("toTabAligned: %w", err)
				return
			}
		}
	}()

	return out
}

package main

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

type Holiday struct {
	Date time.Time
	Name string
}

func (me Holiday) String() string {
	return fmt.Sprintf("%s (%s)", me.Date.Format("2006-01-02"), me.Name)
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	//
	// 標準入力から読み取ったデータを 国民の祝日 CSV ファイル として解釈し
	// 本日以降の祝日を出力する.
	//
	// (*) 国民の祝日 CSV ファイルは、エンコーディングが Shift-JIS となっている
	// (*) CSVファイルのURLは、Taskfile.yml を参照
	//

	// Reader
	var (
		bufReader  = bufio.NewReader(os.Stdin)
		decoder    = japanese.ShiftJIS.NewDecoder()
		sjisReader = transform.NewReader(bufReader, decoder)
		csvReader  = csv.NewReader(sjisReader)
	)

	// Channels
	var (
		dataCh   = make(chan Holiday)
		resultCh = make(chan Holiday)
		errCh    = make(chan error)
	)
	defer close(errCh)

	// 読み込みとフィルタリング開始
	go read(csvReader, dataCh, errCh)
	go filter(dataCh, today(), resultCh, errCh)

	// 結果出力
LOOP:
	for {
		select {
		case err := <-errCh:
			return err
		case result, ok := <-resultCh:
			if !ok {
				break LOOP
			}
			fmt.Println(result)
		}
	}

	return nil
}

func today() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
}

func read(in *csv.Reader, outCh chan<- Holiday, errCh chan<- error) {
	defer close(outCh)

	// 一行目はヘッダなので読み飛ばし
	_, err := in.Read()
	if err != nil {
		errCh <- err
		return
	}

	for {
		row, err := in.Read()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			errCh <- err
			break
		}

		// 日付のフォーマットは yyyy/M/d となっている
		t, err := time.Parse("2006/1/2", row[0])
		if err != nil {
			errCh <- err
			break
		}

		outCh <- Holiday{Date: t, Name: row[1]}
	}
}

func filter(inCh <-chan Holiday, sentinel time.Time, outCh chan<- Holiday, errCh chan<- error) {
	defer close(outCh)

	for v := range inCh {
		if v.Date.Compare(sentinel) < 0 {
			continue
		}

		outCh <- v
	}
}

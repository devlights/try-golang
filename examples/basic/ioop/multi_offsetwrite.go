package ioop

import (
	"encoding/hex"
	"errors"
	"io"
	"os"
	"sync"
)

// MultiOffsetWrite は、io.OffsetWriter を非同期で複数実行し、それぞれ異なるオフセット位置に書き込むサンプルです.
//
// REFERENCES:
//   - https://pkg.go.dev/io@go1.25.4#OffsetWriter
func MultiOffsetWrite() error {
	var (
		file *os.File
		err  error
	)
	file, err = os.CreateTemp(os.TempDir(), "trygolang")
	if err != nil {
		return err
	}
	defer os.Remove(file.Name())

	//
	// ファイル内は指定バイト数ごとに区画が決まっているとする
	// 各区画は独立しており、それぞれデータ列を持つとする
	// 以下は、各区画ごとに *io.OffsetWriter を割り当てて一気に書き込んでいる
	//
	const (
		NUM_REGIONS = 10 // 区画の数
	)
	var (
		errCh = make(chan error, NUM_REGIONS)
		wg    sync.WaitGroup
	)
	for i := range NUM_REGIONS {
		var (
			ow = io.NewOffsetWriter(file, int64(i*20))
		)
		wg.Go(func() {
			var (
				buf = []byte("1234567890")
			)
			if _, err := ow.Write(buf); err != nil {
				errCh <- err
			}
		})
	}

	wg.Wait()
	close(errCh)

	var (
		sErrs = make([]error, 0)
	)
	for err := range errCh {
		if err != nil {
			sErrs = append(sErrs, err)
		}
	}
	if len(sErrs) > 0 {
		return errors.Join(sErrs...)
	}

	//
	// 読み直し
	//
	var (
		data   []byte
		dumper = hex.Dumper(os.Stdout)
	)
	defer dumper.Close()

	if data, err = os.ReadFile(file.Name()); err != nil {
		return err
	}

	dumper.Write(data)

	return nil
}

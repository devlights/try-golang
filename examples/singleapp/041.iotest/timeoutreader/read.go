package timeoutreader

import (
	"errors"
	"io"
	"time"
)

var (
	interval     = 10 * time.Millisecond
	ErrRetryOver = errors.New("retry over")
)

func readAllAtOnce(r io.Reader, p []byte) error {
	var (
		b   []byte
		err error
	)

	b, err = io.ReadAll(r)
	if err != nil {
		return err
	}

	copy(p, b)

	return nil
}

func readWithRetry(r io.Reader, p []byte, retries int) error {
	var (
		buf      = make([]byte, 1<<9)
		numRead  int
		offset   int
		count    int
		maxCount = retries + 1
		err      error
	)

	for count = 0; count < maxCount; {
		if len(p) <= offset {
			break
		}

		clear(buf)

		numRead, err = r.Read(buf)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			time.Sleep(interval)
			count++
			continue
		}

		copy(p[offset:offset+numRead], buf[:numRead])
		offset += numRead
	}

	if maxCount <= count {
		return ErrRetryOver
	}

	return nil
}

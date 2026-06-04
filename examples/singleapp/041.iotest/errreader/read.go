package errreader

import (
	"errors"
	"io"
	"syscall"
	"time"
)

var (
	Interval         = 10 * time.Millisecond
	ErrTooManyEAGAIN = errors.New("retry over (EAGAIN)")
	MaxRetryCount    = 5
)

func read(r io.Reader, p []byte) error {
	var (
		buf     = make([]byte, 2)
		numRead int
		offset  int
		count   int
		err     error
	)

	for count = 0; count < MaxRetryCount; {
		if len(p) <= offset {
			break
		}

		clear(buf)

		numRead, err = r.Read(buf)
		if err != nil {
			switch {
			case errors.Is(err, io.EOF):
				break
			case errors.Is(err, syscall.EAGAIN):
				time.Sleep(Interval)
				count++
				continue
			default:
				return err
			}
		}

		copy(p[offset:offset+numRead], buf[:numRead])
		offset += numRead
	}

	if MaxRetryCount <= count {
		return ErrTooManyEAGAIN
	}

	return nil
}

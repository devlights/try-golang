package main

import (
	"context"
	"errors"
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type (
	// rot13Reader は、ROT13エンコーディングを行います。
	//
	// REFERENCES:
	// 	- https://ja.wikipedia.org/wiki/ROT13
	rot13Reader struct {
		ctx context.Context
		r   io.Reader
	}
)

func NewRot13(ctx context.Context, r io.Reader) *rot13Reader {
	return &rot13Reader{ctx, r}
}

func (me *rot13Reader) Read(p []byte) (int, error) {
	select {
	case <-me.ctx.Done():
		return 0, me.ctx.Err()
	default:
	}

	n, err := me.r.Read(p)
	if err != nil {
		return n, err
	}

	if n > 0 {
		for i := range n {
			p[i] = me.toRot13(p[i])
		}
	}

	select {
	case <-me.ctx.Done():
		return 0, me.ctx.Err()
	default:
		return n, err
	}
}

func (me *rot13Reader) toRot13(b byte) byte {
	switch {
	case 'A' <= b && b <= 'Z':
		return (b-'A'+13)%26 + 'A'
	case 'a' <= b && b <= 'z':
		return (b-'a'+13)%26 + 'a'
	default:
		return b
	}
}

func main() {
	log.SetFlags(log.Lmicroseconds)

	var (
		rootCtx  = context.Background()
		ctx, cxl = signal.NotifyContext(rootCtx, syscall.SIGINT)
		err      error
	)
	defer cxl()

	if err = run(ctx); err != nil {
		if errors.Is(err, context.Canceled) {
			return
		}

		log.Fatal(err)
	}
}

func run(pCtx context.Context) error {
	var (
		r   = NewRot13(pCtx, os.Stdin)
		err error
	)
	if _, err = io.Copy(os.Stdout, r); err != nil {
		return err
	}

	return nil
}

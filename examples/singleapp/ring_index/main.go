package main

import (
	"context"
	"fmt"
	"log"
)

type (
	ring struct {
		curr int
		last int
		next int
	}
)

func newRing() *ring {
	return &ring{curr: 0, last: 2, next: 1}
}

func (me *ring) rotate() {
	me.curr = (me.curr + 1) % 3
	me.last = (me.last + 1) % 3
	me.next = (me.next + 1) % 3
}

func (me *ring) String() string {
	return fmt.Sprintf("(curr=%d,last=%d,next=%d)", me.curr, me.last, me.next)
}

func main() {
	log.SetFlags(0)

	ctx := context.Background()
	if err := run(ctx); err != nil {
		log.Panic(err)
	}
}

func run(ctx context.Context) error {
	var (
		r = newRing()
	)

	for range 10 {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		log.Println(r)
		r.rotate()
	}

	return nil
}

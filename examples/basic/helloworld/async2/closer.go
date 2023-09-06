package async2

import "sync"

type _closer struct {
	wg *sync.WaitGroup
	ch chan<- string
}

func newCloser(wg *sync.WaitGroup, ch chan<- string) *_closer {
	c := new(_closer)

	c.wg = wg
	c.ch = ch

	return c
}

func (me *_closer) run() {
	defer close(me.ch)
	me.wg.Wait()
}

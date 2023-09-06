package async2

import "log"

type _printer struct {
	ch <-chan string
}

func newPrinter(ch <-chan string) *_printer {
	p := new(_printer)
	p.ch = ch
	return p
}

func (me *_printer) run() {
	for v := range me.ch {
		log.Println(v)
	}
}

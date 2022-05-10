package index

import (
	"log"
	"net/http"
)

type Handler []byte

// Implements
var _ http.Handler = (Handler)(nil)

func (me Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write(me); err != nil {
		log.Fatal(err)
	}
}

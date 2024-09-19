//go:build linux

package main

import (
	"log"

	"github.com/devlights/try-golang/examples/singleapp/dev_shm/shm"
)

func main() {
	log.SetFlags(0)

	if err := run(); err != nil {
		log.Panic(err)
	}
}

func run() error {
	m, err := shm.New("test1")
	if err != nil {
		return err
	}
	defer m.Close()

	if _, err = m.Write([]byte("hello")); err != nil {
		return err
	}

	return nil
}

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
	m, err := shm.Open("test1")
	if err != nil {
		return err
	}
	defer m.Close()

	buf := make([]byte, 1<<6)
	n, err := m.Read(buf)
	if err != nil {
		return err
	}

	log.Println(n, string(buf[:n]))

	return nil
}

package main

import (
	"errors"
	"io"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	buf := make([]byte, 1)
	for {
		clear(buf)
		_, err := conn.Read(buf)
		if errors.Is(err, io.EOF) {
			return
		}
	}
}

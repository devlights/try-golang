package main

import (
	"bytes"
	"errors"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	log.SetFlags(log.Lmicroseconds)

	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	conn, err := net.Dial("tcp", ":8888")
	if err != nil {
		return err
	}
	defer func() {
		conn.Close()
		log.Println("[TCP-C] close")
	}()
	log.Println("[TCP-C] connect tcp-server")

	buf := make([]byte, 5)
	n, err := conn.Read(buf)
	if err != nil {
		switch {
		case errors.Is(err, io.EOF):
			return nil
		default:
			return err
		}
	}

	msg := buf[:n]
	log.Printf("[TCP-C] recv (%s)", msg)

	msg = bytes.ToUpper(msg)
	_, err = conn.Write(msg)
	if err != nil {
		return err
	}
	log.Printf("[TCP-C] send (%s)", msg)

	buf = make([]byte, 1)
	for {
		conn.SetReadDeadline(time.Now().Add(100 * time.Millisecond))

		_, err = conn.Read(buf)
		if err != nil {
			if errors.Is(err, io.EOF) {
				log.Println("[TCP-C] disconnect")
				break
			}

			return err
		}
	}

	return nil
}

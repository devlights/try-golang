/*
UNIXドメインソケット 抽象名前空間のサンプル（netパッケージのConnを利用する版)
*/
package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"time"
)

type (
	Args struct {
		IsServer bool
	}
)

const (
	// 抽象名前空間のソケットアドレス（@記号で始まる名前は\0に変換される）
	serverAddr = "@go_unix_domain_socket_test"
	bufSize    = len("hello")
)

var (
	args Args
)

func init() {
	flag.BoolVar(&args.IsServer, "server", false, "server mode")
}

func main() {
	log.SetFlags(log.Lmicroseconds)
	flag.Parse()

	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	var err error
	switch args.IsServer {
	case true:
		err = runServer()
	default:
		err = runClient()
	}

	if err != nil {
		return err
	}

	return nil
}

func runServer() error {
	ln, err := net.Listen("unix", serverAddr)
	if err != nil {
		return err
	}
	defer ln.Close()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)
	go func(sigCh <-chan os.Signal) {
		<-sigCh
		log.Println("[S] Shutdown...")
		ln.Close()
		os.Exit(0)
	}(sigCh)

	buf := make([]byte, bufSize)
	for {
		conn, err := ln.Accept()
		if err != nil {
			if errors.Is(err, net.ErrClosed) {
				return nil
			}

			return err
		}

		// サンプルなので１接続で占有状態とする
		errCh := make(chan error)
		go func() {
			defer close(errCh)
			defer func() {
				conn.Close()
				log.Println("[S] close")
			}()

			clear(buf)
			n, err := conn.Read(buf)
			if n == 0 || errors.Is(err, io.EOF) {
				log.Println("[S] disconnect")
				return
			}

			if err != nil {
				errCh <- err
			}

			message := buf[:n]
			log.Printf("[S] Recv (%s)", message)

			message = bytes.ToUpper(buf[:n])
			_, err = conn.Write(message)
			if err != nil {
				errCh <- err
			}
			log.Printf("[S] Send (%s)", message)

			// FIN待機用のタイムアウト設定
			err = conn.SetReadDeadline(time.Now().Add(2 * time.Second))
			if err != nil {
				log.Printf("[S] SetReadDeadline error: %v", err)
				errCh <- err
				return
			}

			for {
				clear(buf)
				if n, err = conn.Read(buf); n == 0 || errors.Is(err, io.EOF) {
					log.Println("[S] disconnect")
					break
				}

				if err != nil {
					if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
						log.Println("[S] timeout (client FIN)")
						break
					}

					errCh <- err
					return
				}
			}
		}()

		err = <-errCh
		if err != nil {
			return err
		}
	}
}

func runClient() error {
	conn, err := net.Dial("unix", serverAddr)
	if err != nil {
		return err
	}
	defer func() {
		conn.Close()
		log.Println("[C] close")
	}()

	buf := make([]byte, bufSize)
	copy(buf, []byte("hello"))

	_, err = conn.Write(buf)
	if err != nil {
		return err
	}
	log.Printf("[C] Send (%s)", buf)

	clear(buf)
	n, err := conn.Read(buf)
	if n == 0 || errors.Is(err, io.EOF) {
		log.Println("[S] disconnect")
		return nil
	}
	if err != nil {
		return err
	}

	message := string(buf[:n])
	log.Printf("[C] Recv (%s)", message)

	// Graceful shutdown
	{
		unixConn, ok := conn.(*net.UnixConn)
		if !ok {
			return fmt.Errorf("conn.(*net.UnixConn) failed")

		}

		err = unixConn.CloseWrite()
		if err != nil {
			return err
		}
		log.Println("[C] SEND FIN (shutdown(SHUT_WR))")

		for {
			clear(buf)
			if n, err = conn.Read(buf); n == 0 || errors.Is(err, io.EOF) {
				log.Println("[C] disconnect")
				break
			}
		}
	}

	return nil
}

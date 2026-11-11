package main

import (
	"context"
	"errors"
	"log"
	"net"
	"os/signal"
	"syscall"
	"time"
)

const (
	Addr = "127.0.0.1"
	Port = "9999"
)

func main() {
	var (
		ctx, cxl = signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
		err      error
	)
	defer cxl()

	if err = run(ctx); err != nil {
		panic(err)
	}
}

func run(pCtx context.Context) error {
	var (
		ctx, cxl = context.WithCancel(pCtx)
		conn     net.PacketConn
		err      error
	)
	defer cxl()

	if conn, err = net.ListenPacket("udp4", net.JoinHostPort(Addr, Port)); err != nil {
		return err
	}
	defer func() {
		conn.Close()
	}()

	log.Printf("[recv] listening on %s", conn.LocalAddr())

LOOP:
	for {
		select {
		case <-ctx.Done():
			log.Println("[recv] shutdown...")
			break LOOP
		default:
		}

		var (
			n    int
			from net.Addr
			err  error

			buf = make([]byte, 4096)
		)
		if err = conn.SetReadDeadline(time.Now().Add(1 * time.Second)); err != nil {
			return err
		}

		if n, from, err = conn.ReadFrom(buf); err != nil {
			var (
				ne, ok = errors.AsType[net.Error](err)
			)
			if !ok || !ne.Timeout() {
				return err
			}

			continue
		}

		log.Printf("[recv] from=%s %q", from, buf[:n])
	}

	return nil
}

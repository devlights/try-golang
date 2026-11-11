package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/net/ipv4"
)

const (
	Addr = "127.0.0.1"
	Port = "0"
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
	defer conn.Close()

	//
	// net.Connから*ipv4.PacketConnを作成してAF11パケットを設定する。
	// *ipc4.PacketConn は、IPヘッダのオプション値を変更したり出来る。
	//
	var (
		p = ipv4.NewPacketConn(conn)
	)

	//
	// AF11パケットとしてTOS設定
	//
	// AF11パケットとは、IPヘッダーのDSCPに「AF11」という値を設定したパケットのこと。
	// QoSの分類の一つで、「AFクラス1に属し、そのクラス内では破棄されにくい」という扱いを表す。
	//
	// DSCP値は10進数で 10、2進数で 001010 となる。
	// 値としては、TOSは１バイトであり、現在の用途では上位6ビットがDSCPで、下位2ビットがECNとなる。
	// ECNは通常0なので、AF11のTOSの値は DSCP=10(0b001010), ECN=0(00)で、0b00101000 となる。
	//
	// なので、 0b00101000 (bin) = 0x28 (hex) = 40 (dec) となる。
	//
	const (
		dscpAF11 = 10
		tosAF11  = dscpAF11 << 2 // 0x28：DSCP=AF11、ECN=00
	)
	if err = p.SetTOS(tosAF11); err != nil {
		return err
	}

	//
	// 送信
	//
	var (
		dst = &net.UDPAddr{
			IP:   net.IPv4(127, 0, 0, 1),
			Port: 5000,
		}
	)

	log.Printf(
		"[send] from=%s to=%s DSCP=%d TOS=0x%02x",
		conn.LocalAddr(),
		dst,
		dscpAF11,
		tosAF11,
	)

	for i := 0; i < 5; i++ {
		select {
		case <-ctx.Done():
			return nil
		default:
			var (
				m                        = fmt.Appendf(nil, "hello AF11: seq=%d", i)
				cm  *ipv4.ControlMessage = nil
				n   int
				err error
			)
			if err = p.SetWriteDeadline(time.Now().Add(500 * time.Millisecond)); err != nil {
				return err
			}

			if n, err = p.WriteTo(m, cm, dst); err != nil {
				return err
			}

			log.Printf("[send] sent bytes=%d message=%q", n, string(m))
		}

		if i < 5 {
			time.Sleep(500 * time.Millisecond)
		}
	}

	return nil
}

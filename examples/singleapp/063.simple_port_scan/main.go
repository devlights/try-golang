package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"
)

type Args struct {
	host string
	from int
	to   int
}

var (
	debug = flag.Bool("debug", false, "debug")
	args  = Args{}
)

func init() {
	log.SetFlags(0)
}

func panicErr[T any](o T, err error) T {
	if err != nil {
		log.Panic(err)
	}

	return o
}

func main() {
	flag.Parse()

	if flag.NArg() != 3 {
		log.Println("Usage: simple_port_scan host port(from) port(to)")
		os.Exit(1)
	}

	a := flag.Args()
	args.host = a[0]
	args.from = panicErr(strconv.Atoi(a[1]))
	args.to = panicErr(strconv.Atoi(a[2]))

	var l net.Listener
	var e error
	if *debug {
		l, e = net.Listen("tcp4", fmt.Sprintf("127.0.0.1:%d", ((args.to-args.from)/2)+args.from))
		if e != nil {
			log.Println(e)
		}
	}
	defer l.Close()

	if err := run(); err != nil {
		log.Panic(err)
	}
}

func run() error {
	for i := args.from; i <= args.to; i++ {
		var (
			addr = fmt.Sprintf("%s:%d", args.host, i)
			conn net.Conn
			err  error
		)

		conn, err = net.DialTimeout("tcp4", addr, 1*time.Second)
		switch err {
		case nil:
			log.Printf("%d open", i)
			conn.Close()
		default:
			log.Printf("%d closed", i)
		}
	}

	return nil
}

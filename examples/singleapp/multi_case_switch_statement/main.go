package main

import (
	"log"
	"math/rand"
	"time"
)

var (
	rnd *rand.Rand
)

func init() {
	log.SetFlags(0)
	rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func main() {
	log.SetFlags(0)

	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	var (
		v = (1 + rnd.Intn(10))
	)

	log.Printf("v=%d", v)

	//
	// Goの switch statement は case に複数の値を指定出来る
	// 複数指定する場合はカンマで並べる。
	//
	// REFERENCES:
	//   - https://go.dev/ref/spec#Switch_statements
	//   - https://www.w3schools.com/go/go_switch_multi.php
	//
	switch v {
	case 1, 2, 3:
		log.Printf("p1")
	case 4, 5:
		log.Printf("p2")
	case 6, 7, 8:
		log.Printf("p3")
	default:
		log.Printf("p4")
	}

	return nil
}

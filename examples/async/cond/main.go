package main

import (
	"flag"
	"log"
	"math/rand/v2"
	"os"
	"os/signal"
	"sync"
	"time"
)

const (
	MAX_ITEM_COUNT = 10
	MAX_PRODUCERS  = 2
	MAX_CONSUMERS  = MAX_ITEM_COUNT / 2
	WATCH_INTERVAL = 1 * time.Second
)

type (
	Args struct {
		debug bool
	}
)

var (
	args Args
)

func init() {
	flag.BoolVar(&args.debug, "debug", false, "debug mode")
}

func main() {
	//
	// Producer-Consumer-Watcher のサンプル
	// 各非同期処理のランデブーポイントの制御に
	// *sync.Cond を利用。
	//

	log.SetFlags(log.Lmicroseconds)
	log.SetOutput(os.Stdout)

	flag.Parse()

	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	var (
		ch   = make(chan int, MAX_ITEM_COUNT*MAX_PRODUCERS)
		sig  = make(chan os.Signal, 1)
		done = make(chan struct{})
	)
	defer close(ch)

	signal.Notify(sig, os.Interrupt)
	go func() {
		<-sig
		log.Printf("<<Interrupt>>")
		close(done)
	}()

	var (
		producer = sync.NewCond(&sync.Mutex{})
		consumer = sync.NewCond(&sync.Mutex{})
	)

	// 役割： 生産者と消費者を監視し、必要であれば叩き起こす
	go watch(1, ch, producer, consumer)

	// 役割： 生産を行う
	for i := range MAX_PRODUCERS {
		go produce(i+1, ch, (i+1)*10000, producer, consumer)
	}

	// 役割： 消費を行う
	for i := range MAX_CONSUMERS {
		go consume(i+1, ch, consumer)
	}

	// ゴルーチンの終了待機などについては割愛
	<-done
	log.Printf("<<DONE>>")

	return nil
}

func watch(id int, ch <-chan int, producer, consumer *sync.Cond) {
	for {
		func() {
			producer.L.Lock()
			defer producer.L.Unlock()

			if len(ch) == 0 {
				producer.Broadcast()
				log.Printf("[W][%02d] >>> 0個です。生産しなさい。", id)
			}
		}()

		<-time.After(WATCH_INTERVAL)

		func() {
			consumer.L.Lock()
			defer consumer.L.Unlock()

			if len(ch) != 0 {
				consumer.Broadcast()
				log.Printf("[W][%02d] >>> 生産されています。消費しなさい。", id)
			}
		}()
	}
}

func produce(id int, ch chan<- int, start int, producer, consumer *sync.Cond) {
	var (
		count int
	)
	for i := start; ; {
		func() {
			producer.L.Lock()
			defer producer.L.Unlock()

			for len(ch) > cap(ch)/2 {
				dbg("[P][%02d] <<< 消費されるまで待機します。(残:%d)", id, len(ch))
				producer.Wait()
			}
		}()

		func() {
			consumer.L.Lock()
			defer consumer.L.Unlock()

			count = rand.IntN(MAX_ITEM_COUNT)
			for c := range count {
				ch <- i + (c + 1)
			}

			log.Printf("[P][%02d] >>> %d個生産しました。(残:%d)", id, count, len(ch))
			consumer.Broadcast()

			i += count
		}()

		// 次のタスク着手まで少し休憩
		<-time.After(time.Duration(rand.IntN(500)) * time.Millisecond)
	}
}

func consume(id int, ch <-chan int, consumer *sync.Cond) {
	for {
		func() {
			consumer.L.Lock()
			defer consumer.L.Unlock()

			for len(ch) == 0 {
				dbg("[C][%02d] <<< 生産されるまで待機します。(残:%d)", id, len(ch))
				consumer.Wait()
			}

			log.Printf("[C][%02d] >>> 消費しました (%v)(残:%d)", id, <-ch, len(ch))
		}()

		// 次のタスク着手まで少し休憩
		<-time.After(time.Duration(rand.IntN(1000)) * time.Millisecond)
	}
}

func dbg(format string, v ...any) {
	if args.debug {
		log.Printf(format, v...)
	}
}

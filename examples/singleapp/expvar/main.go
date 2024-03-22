package main

import (
	"expvar"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

type _dbg struct {
	counter *expvar.Int
	message *expvar.String
	values  *expvar.Map
}

func newDbg() *_dbg {
	return &_dbg{
		counter: expvar.NewInt("counter"),
		message: expvar.NewString("message"),
		values:  expvar.NewMap("values"),
	}
}

var (
	dbg = newDbg()
)

func main() {
	go func() {
		// localhost:8888/debug/vars で値が表示されるようになる
		// デフォルトで "cmdline" と "memstats" というキーが公開される
		log.Fatal(http.ListenAndServe(":8888", nil))
	}()

	go func() {
		for {
			<-time.After(1 * time.Second)

			//
			// expvar で公開設定している各変数を更新
			// (ブラウザなどで見ると反映されていることが分かる)
			//
			dbg.counter.Add(1)
			dbg.message.Set(fmt.Sprintf("hello-%02d", dbg.counter.Value()))

			// expvar.NewInt()を呼ぶと内部でPublish()が呼ばれてしまうので、直接生成する
			c := &expvar.Int{}
			c.Set(dbg.counter.Value() * dbg.counter.Value())
			dbg.values.Set("counter", c)

			// expvar.NewString()を呼ぶと内部でPublish()が呼ばれてしまうので、直接生成する
			s := &expvar.String{}
			s.Set(strings.ToUpper(dbg.message.Value()))
			dbg.values.Set("message", s)
		}
	}()

	for {
		<-time.After(1 * time.Minute)
	}
}

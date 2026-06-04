package main

import (
	"log"
	"sync"
	"time"

	"golang.org/x/sync/singleflight"
)

func init() {
	log.SetFlags(0)
}

func heavy(delay time.Duration, prefix string) {
	log.Printf("%s start", prefix)
	defer log.Printf("%s end  ", prefix)

	time.Sleep(delay)
}

func main() {
	// Group.Do() or Group.DoChan() を利用して
	// 重複する呼び出しが発生する箇所や処理負荷が高い操作などの呼び出しを
	// 抑制することが出来る。
	//
	// Cache Stampedeが発生する可能性がある部分には非常に有効です。

	const (
		KEY = "FUNC-GROUP-KEY"
	)

	var (
		grp     = &singleflight.Group{}
		ready   = make(chan struct{})
		results = make(chan (<-chan singleflight.Result))
		wg      = sync.WaitGroup{}
	)

	wg.Add(3)
	go func() {
		defer wg.Done()

		results <- grp.DoChan(KEY, func() (any, error) {
			<-ready
			heavy(3*time.Second, "func1")
			return "func1", nil
		})
	}()
	go func() {
		defer wg.Done()

		results <- grp.DoChan(KEY, func() (any, error) {
			<-ready
			heavy(1*time.Second, "func2")
			return "func2", nil
		})
	}()
	go func() {
		defer wg.Done()

		results <- grp.DoChan(KEY, func() (any, error) {
			<-ready
			heavy(2*time.Second, "func3")
			return "func3", nil
		})
	}()
	go func() {
		defer close(results)
		wg.Wait()
	}()

	// よーいドン
	close(ready)

	for ret := range results {
		log.Printf("%+v", <-ret)
	}

	// 更に追加呼び出し
	//   ただし、実行する前に Group.Forget() を呼び出して
	//   キーに紐づく結果を忘れさせてから実行
	grp.Forget(KEY)
	ret := grp.DoChan(KEY, func() (any, error) {
		heavy(1*time.Second, "func4")
		return "func4", nil
	})
	log.Printf("%+v", <-ret)

	/*
	   task: [build] go build -o app
	   task: [run] ./app
	   func2 start
	   func2 end
	   {Val:func2 Err:<nil> Shared:true}
	   {Val:func2 Err:<nil> Shared:true}
	   {Val:func2 Err:<nil> Shared:true}
	   func4 start
	   func4 end
	   {Val:func4 Err:<nil> Shared:false}
	*/
}

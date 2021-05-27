package gotour29

import (
	"fmt"
	"sync"
	"time"
)

type (
	// Incrementer -- 何かを加算する振る舞いを持ちます.
	Incrementer interface {
		// Increment -- 加算します.
		Increment(wg *sync.WaitGroup)
	}
)

type (
	counter struct {
		Val int
	}

	// NotSafeCounter -- 安全ではないカウンター
	NotSafeCounter struct {
		counter
	}

	// SafeCounter -- 安全なカウンター
	SafeCounter struct {
		counter
		mux sync.Mutex
	}
)

func (c *counter) String() string {
	return fmt.Sprintf("Val: %d", c.Val)
}

// NewNotSafeCounter -- 安全ではないカウンターを返します.
func NewNotSafeCounter() Incrementer {
	return &NotSafeCounter{
		counter{Val: 0},
	}
}

// NewSafeCounter -- 安全なカウンターを返します.
func NewSafeCounter() Incrementer {
	return &SafeCounter{
		counter: counter{Val: 0},
	}
}

// Increment -- 加算します.
func (c *NotSafeCounter) Increment(wg *sync.WaitGroup) {
	defer wg.Done()

	cur := c.Val
	time.Sleep(1 * time.Microsecond)
	cur++
	time.Sleep(1 * time.Microsecond)
	c.Val = cur
}

// Increment -- 加算します.
func (c *SafeCounter) Increment(wg *sync.WaitGroup) {
	defer wg.Done()

	// 排他制御
	c.mux.Lock()
	defer c.mux.Unlock()

	cur := c.Val
	time.Sleep(1 * time.Microsecond)
	cur++
	time.Sleep(1 * time.Microsecond)
	c.Val = cur
}

// Mutex は、 Tour of Go - sync.Mutex (https://tour.golang.org/concurrency/9) の サンプルです。
func Mutex() error {
	// ------------------------------------------------------------
	// Mutex
	//
	// 基本的に他の言語のmutexと考え方も使い方も同じ.
	// クリティカルセクションになる部分で、 Lock() と Unlock() を呼び出して排他制御。
	// Unlock() は、エラー発生時などで呼べない可能性を考慮して defer で呼ぶのが多い。
	// ------------------------------------------------------------
	var (
		wg1     sync.WaitGroup
		wg2     sync.WaitGroup
		notSafe = NewNotSafeCounter()
		safe    = NewSafeCounter()
		times   = 100
	)

	wg1.Add(times)
	increment(notSafe, times, &wg1)
	wg1.Wait()

	fmt.Printf("[NotSafeCounter] %v\n", notSafe)

	wg2.Add(times)
	increment(safe, times, &wg2)
	wg2.Wait()

	fmt.Printf("[SafeCounter] %v\n", safe)

	return nil
}

func increment(inc Incrementer, times int, wg *sync.WaitGroup) {
	for i := 0; i < times; i++ {
		go inc.Increment(wg)
	}
}

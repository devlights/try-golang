package main

import (
	"fmt"
	"io"
	"sync"
	"testing"
)

func BenchmarkChanSync(b *testing.B) {
	benchcases := []struct {
		Name string
		In   int
	}{
		{"100", 100},
		{"1000", 1000},
		{"10000", 10000},
		{"100000", 100000},
		{"1000000", 1000000},
	}

	for _, bc := range benchcases {
		bc := bc
		b.Run(bc.Name, func(b *testing.B) {
			for b.Loop() {
				var (
					ch = make(chan int, bc.In)
				)
				for i := range bc.In {
					ch <- i
				}
				close(ch)

				var (
					total uint64
				)
				for v := range ch {
					total += uint64(v)
				}

				fmt.Fprintln(io.Discard, total)
			}
		})
	}
}

func BenchmarkChanAsync(b *testing.B) {
	benchcases := []struct {
		Name string
		In   int
	}{
		{"100", 100},
		{"1000", 1000},
		{"10000", 10000},
		{"100000", 100000},
		{"1000000", 1000000},
	}

	for _, bc := range benchcases {
		bc := bc
		b.Run(bc.Name, func(b *testing.B) {
			for b.Loop() {
				var (
					ch = make(chan int, bc.In)
					wg = sync.WaitGroup{}
				)
				wg.Add(2)

				go func(ch chan<- int) {
					defer wg.Done()
					defer close(ch)

					for i := range bc.In {
						ch <- i
					}
				}(ch)

				go func(ch <-chan int) {
					defer wg.Done()

					var (
						total uint64
					)
					for v := range ch {
						total += uint64(v)
					}

					fmt.Fprintln(io.Discard, total)
				}(ch)

				wg.Wait()
			}
		})
	}
}

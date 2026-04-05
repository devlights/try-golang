package main

import "fmt"

func main() {
	// 処理に特に意味は無い

	var (
		p = func() <-chan int {
			out := make(chan int)
			go func() {
				defer close(out)
				for i := range 5 {
					out <- i
				}
			}()
			return out
		}
		c = func(in <-chan int) <-chan bool {
			out := make(chan bool)
			go func() {
				defer close(out)
				for v := range in {
					fmt.Println(v)
				}
			}()
			return out
		}
	)
	<-c(p())
}

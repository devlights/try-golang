package main

func main() {
	ch := make(chan struct{}, 1)
	go func() {
		read(ch)
	}()
	write(ch)
}

func read(ch <-chan struct{}) {
	for {
		<-ch
	}
}

func write(ch chan<- struct{}) {
	for {
		ch <- struct{}{}
	}
}

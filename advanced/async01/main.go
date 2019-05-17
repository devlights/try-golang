package main

import (
	"fmt"
	"time"
)

func main() {
	// タイムアウト待ちをするチャネル
	timeoutChannel := make(chan bool)
	go timeout(timeoutChannel)

	// 入力を読み取るチャネル
	readChannel := make(chan string)
	go readWord(readChannel)

	// メッセージ待機
	select {
	case word := <-readChannel:
		fmt.Println("recv", word)
	case <-timeoutChannel:
		fmt.Println("timeout")
		break
	}
}

func readWord(ch chan string) {
	fmt.Println("Input: ")

	var word string
	_, err := fmt.Scanf("%s", &word)
	if err != nil {
		panic(err.Error())
	}

	ch <- word
}

func timeout(ch chan bool) {
	time.Sleep(5 * time.Second)
	ch <- true
}

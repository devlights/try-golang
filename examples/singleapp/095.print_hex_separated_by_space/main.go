package main

import (
	"crypto/rand"
	"fmt"
)

func main() {
	// fmt.Printf()で指定できるverbで %x , %X の場合
	// '% x' , '% X' と半角スペースを前に付与することで
	// 16進数を1バイトずつスペースで区切って出力することが出来る。
	//
	// https://pkg.go.dev/fmt に以下の記載がある。(Other flags: の部分）
	//
	//    (space) leave a space for elided sign in numbers (% d);
	//    put spaces between bytes printing strings or slices in hex (% x, % X)
	//
	//    空白フラグは数値では省略された符号の位置に空白を置く。
	//    文字列またはスライスを16進数で表示する % x / % X では、バイト間に空白を置く。
	var (
		b = make([]byte, 16)
	)
	_, _ = rand.Read(b) // docには It never returns an error とあり、本APIはエラー処理不要

	fmt.Printf("%x\n", b)
	fmt.Printf("% x\n", b)
	fmt.Printf("% X\n", b)
}

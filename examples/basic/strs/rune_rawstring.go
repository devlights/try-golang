package strs

import "fmt"

// RuneRawString は、Go言語における 文字と文字列とRaw文字列についてのサンプルです
func RuneRawString() error {
	// string. 文字列
	string1 := "こんにちは"

	// Rune. Go言語では文字を表す
	rune1 := []rune(string1)

	// raw string. 生の文字列。Pythonのトリプルクォートと同じ.
	raw1 := `aaaa \nbbbb`

	fmt.Printf("%T, %T, %T\n", rune1, string1, raw1)
	fmt.Printf("%#U, %v, %v\n", rune1, string1, raw1)

	// 文字列を rune の スライスに変換
	string2 := "あいう"
	rune2 := []rune(string2)
	fmt.Println(rune2)

	// 文字列を byte の スライスに変換
	bytes1 := []byte(string2)
	fmt.Println(bytes1)

	// rune は文字数、byte はバイト数
	fmt.Printf("len(rune2): %d, len(bytes1): %d\n", len(rune2), len(bytes1))

	// rune の スライス から 文字列 へ
	string3 := string(rune2)
	fmt.Printf("%s, %s\n", string2, string3)

	// byte の スライス から 文字列 へ
	string4 := string(bytes1)
	fmt.Printf("%s, %s\n", string2, string4)

	return nil

	/*
	   $ task
	   task: Task "build" is up to date
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: string_rune_rawstring

	   [Name] "string_rune_rawstring"
	   []int32, string, string
	   [U+3053 'こ' U+3093 'ん' U+306B 'に' U+3061 'ち' U+306F 'は'], こんにちは, aaaa \nbbbb
	   [12354 12356 12358]
	   [227 129 130 227 129 132 227 129 134]
	   len(rune2): 3, len(bytes1): 9
	   あいう, あいう
	   あいう, あいう


	   [Elapsed] 58.851µs
	*/

}

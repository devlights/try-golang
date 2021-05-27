package strs

import "fmt"

// StringRuneRawString は、Go言語における 文字と文字列とRaw文字列についてのサンプルです
func StringRuneRawString() error {
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
}

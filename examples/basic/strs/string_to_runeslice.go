package strs

import "fmt"

// StringToRuneSlice は、文字列とルーンスライスの遷移を表示するサンプルです
func StringToRuneSlice() error {

	// Go の 文字列 は、他の言語と同様に immutable となっている
	// なので、一度作成した文字列を変更することは出来ない
	// 変更したい場合は、新たな文字列を作って格納する必要がある
	s := "hello world"

	// 文字列を runeスライス に変換
	r := []rune(s)

	// 変更
	r[0] = 'H'

	// 再度文字列へ
	s2 := string(r)

	fmt.Printf("Before:\t%q\tAfter:\t%q\tRune:%v\n", s, s2, r)

	return nil
}

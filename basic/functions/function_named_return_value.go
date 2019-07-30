package functions

import "fmt"

func oneNamedReturnValue(s string) (r string) {
	r = fmt.Sprintf("Hello %s", s)

	// result parameter を利用している場合、戻り値はすでに変数として
	// 定義されている状態なので、値をいれたら、そのまま return で良い
	return
}

func multiNamedReturnValue(s string) (original string, r string) {
	original = s
	r = oneNamedReturnValue(s)

	return
}

func FunctionNamedReturnValue() error {

	var (
		s = "world"
	)

	orig, r := multiNamedReturnValue(s)

	fmt.Printf("%s\n", oneNamedReturnValue(s))
	fmt.Printf("orig=%s, result=%s\n", orig, r)

	return nil
}

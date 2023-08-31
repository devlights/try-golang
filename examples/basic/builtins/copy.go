package builtins

import "github.com/devlights/gomy/output"

// Copy は、ビルトイン関数 copy についてのサンプルです.
//
// # REFERENCES
//   - https://pkg.go.dev/builtin@go1.21.0#copy
func Copy() error {
	//
	// copy(dst, src) は、srcからdstへ要素をコピーする
	//   - 特例として dst=[]byte, src=string の場合は動作するようになっている
	//
	var (
		src1 = []int{1, 2, 3, 4, 5}
		src2 = src1[:3]
		dst1 = make([]int, len(src1))
		dst2 = make([]int, len(src1)-1)
		dst3 = make([]int, len(src1)-1)
		p    = func() {
			output.Stdoutl("[src1]", src1)
			output.Stdoutl("[src2]", src2)
			output.Stdoutl("[dst1]", dst1)
			output.Stdoutl("[dst2]", dst2)
			output.Stdoutl("[dst3]", dst3)
			output.StdoutHr()
		}
	)

	p()
	copy(dst1, src1)
	copy(dst2, src1)
	copy(dst3, src2)
	p()
	dst1[0] = 100
	dst2[0] = 200
	dst3[0] = 300
	p()

	var (
		dst4 = make([]byte, 10)
		src3 = "hello world"
	)

	// dst=[]byte, src=string は型が異なるが特別に許可される
	copy(dst4, src3)
	output.Stdoutl("[dst4]", dst4)

	return nil
}

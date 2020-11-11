package sub1

import "github.com/devlights/try-golang/internal/examples/basic/internalpkg/internal/internal/sub2"

// CallSub1 -- サンプルから呼び出される関数
func CallSub1() string {
	return "InternalPkgSub1"
}

// CallSub2 -- サンプルから呼び出される関数
func CallSub2() string {
	return sub2.Sub2()
}

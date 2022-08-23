package args_test

import (
	"strconv"
	"testing"

	args "github.com/devlights/try-golang/examples/singleapp/flag_pkg/unittest"
)

func FuzzArgs(f *testing.F) {
	f.Add(1, "hello")
	f.Add(10, "helloworld")
	f.Add(100, "こんにちは世界")
	f.Add(-1, "こんにちは世界")

	f.Fuzz(func(t *testing.T, i int, s string) {
		options := []string{"-a", strconv.Itoa(i), "-b", s, "-c"}

		sut, err := args.Parse(options)
		if sut == args.Error {
			t.Errorf("i=%v, s=%v, err=%v", i, s, err)
		}
	})
}

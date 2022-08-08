package testpkgcoexist_test

import "testing"

func TestB(t *testing.T) {
	var (
		a, b = 1, 10
	)

	c := a * 10
	if c != b {
		t.Errorf("[want] %v\t[got] %v", b, c)
	}
}

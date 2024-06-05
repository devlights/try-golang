package lib

import "testing"

func TestFn(t *testing.T) {
	const (
		count = 10
	)

	if v := Fn(count); v != count {
		t.Errorf("[want] %v\t[got] %v", count, v)
	}
}

package setupteardown

import (
	"testing"
)

const (
	initX = 0
)

var (
	_x = initX
)

func setupAndTeardown(t *testing.T) func() {
	// Setup
	t.Log("setup")
	_x = 999

	return func() {
		// Teardown
		t.Log("teardown")
		_x = initX
	}
}

func Test_SetupAndTeardown(t *testing.T) {
	defer setupAndTeardown(t)()

	t.Log("test body")
	if _x != 999 {
		t.Errorf("[want] 999\t[got] %v", _x)
	}
}

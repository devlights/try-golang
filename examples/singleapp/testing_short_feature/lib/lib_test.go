package lib_test

import (
	"testing"

	"github.com/devlights/try-golang/cmd/testing_short_feature/lib"
)

func TestAdd(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}

	got := lib.Add(1, 2)
	if got != 3 {
		t.Errorf("[want] 3\t[got] %v\n", got)
	}
}

func TestSum(t *testing.T) {
	if testing.Short() {
		t.Skip("In short mode, this test will be skipped.")
	}

	got := lib.Sum(1, 2, 3, 4, 5)
	if got != 15 {
		t.Errorf("[want] 15\t[got] %v\n", got)
	}
}
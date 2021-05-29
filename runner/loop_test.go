package runner_test

import (
	"bytes"
	"testing"

	"github.com/devlights/try-golang/mapping"
	"github.com/devlights/try-golang/runner"
)

func TestLoop(t *testing.T) {
	r := new(runner.SilentRegister)

	m := mapping.NewSampleMapping()
	m.MakeMapping(r)

	buf := new(bytes.Buffer)
	buf.WriteString("silent\n")

	a := runner.NewLoopArgs(buf, true, m)
	e := runner.NewLoop(a)

	if err := e.Run(); err != nil {
		t.Errorf("should not raise error (%s)", err)
	}

	if !r.Target.Called {
		t.Errorf("never called the target example")
	}
}

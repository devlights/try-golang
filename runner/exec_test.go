package runner_test

import (
	"testing"

	"github.com/devlights/try-golang/mapping"
	"github.com/devlights/try-golang/runner"
)

func TestExec(t *testing.T) {
	r := new(runner.SilentRegister)

	m := mapping.NewSampleMapping()
	m.MakeMapping(r)

	a := runner.NewExecArgs("silent", m)
	e := runner.NewExec(a)

	if err := e.Run(); err != nil {
		t.Errorf("should not raise error (%s)", err)
	}

	if !r.Target.Called {
		t.Errorf("never called the target example")
	}
}

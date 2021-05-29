package mapping_test

import (
	"testing"

	"github.com/devlights/try-golang/mapping"
)

type (
	onlyOneRegister struct{}
)

func (me *onlyOneRegister) Regist(m mapping.ExampleMapping) {
	m["this_is_a_test"] = func() error { return nil }
}

func TestMapping(t *testing.T) {
	m := mapping.NewSampleMapping()
	m.MakeMapping(&onlyOneRegister{})

	_, ok := m["this_is_a_test"]
	if !ok {
		t.Errorf("no key.. [want] false\t[got] true")
	}
}

func TestAllExampleNames(t *testing.T) {
	m := mapping.NewSampleMapping()
	m.MakeMapping(&onlyOneRegister{})

	names := m.AllExampleNames()
	if len(names) != 1 {
		t.Errorf("mismatch example count..\t[want] 1\t[got] %d", len(names))
	}

	name := names[0]
	if name != "this_is_a_test" {
		t.Errorf("mismatch example name..\t[want] this_is_a_test\t[got] %s", name)
	}
}

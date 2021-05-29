package builder_test

import (
	"testing"

	"github.com/devlights/try-golang/builder"
)

func TestMappingIsNotZero(t *testing.T) {
	cases := []struct {
		name string
		out  interface{}
	}{
		{"sizeIsNotZero", 0},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			m := builder.BuildMappings()
			if len(m) == c.out {
				t.Errorf("[want] not zero\t[got] zero")
			}
		})
	}
}

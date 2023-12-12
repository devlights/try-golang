package args_test

import (
	"reflect"
	"strconv"
	"testing"

	args "github.com/devlights/try-golang/examples/singleapp/flag_pkg/unittest"
)

func TestArgs(t *testing.T) {
	tests := []struct {
		in  []string
		out *args.Value
	}{
		{[]string{}, args.Empty},
		{[]string{"-a"}, args.Error},
		{[]string{"-a", "-b"}, args.Error},
		{[]string{"-a", "-b", "-c"}, args.Error},
		{[]string{"-a", "1"}, &args.Value{A: 1}},
		{[]string{"-a", "1", "-b", "hello"}, &args.Value{A: 1, B: "hello"}},
		{[]string{"-a", "1", "-b", "hello", "-c"}, &args.Value{A: 1, B: "hello", C: true}},
		{[]string{"-a", "1", "-b", "hello", "-c", "true"}, &args.Value{A: 1, B: "hello", C: true}},
		{[]string{"-a", "bbb", "-b", "hello", "-c"}, args.Error},
		{[]string{"-a", "bbb", "-b", "hello", "-c", "world"}, args.Error},
	}

	for i, tt := range tests {
		tt := tt
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			sut, err := args.Parse(tt.in)
			if !reflect.DeepEqual(tt.out, sut) {
				t.Errorf("[want] %v\t[got] %v (%v)", tt.out, sut, err)
			}
		})
	}
}

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
		{[]string{"-a", "1"}, &args.Value{1, "", false}},
		{[]string{"-a", "1", "-b", "hello"}, &args.Value{1, "hello", false}},
		{[]string{"-a", "1", "-b", "hello", "-c"}, &args.Value{1, "hello", true}},
		{[]string{"-a", "1", "-b", "hello", "-c", "true"}, &args.Value{1, "hello", true}},
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

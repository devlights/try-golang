package tabledriventesting

import "testing"

func Test_Tdt_01(t *testing.T) {
	tests := []struct {
		input, expected string
	}{
		{"hello", "hello"},
		{"world", "world"},
	}

	for _, tt := range tests {
		tt := tt

		if tt.input != tt.expected {
			t.Errorf("[want] %v\t[got] %v", tt.expected, tt.input)
		}
	}
}

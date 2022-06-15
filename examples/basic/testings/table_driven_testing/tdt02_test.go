package tabledriventesting

import "testing"

func Test_Tdt_02(t *testing.T) {
	tests := []struct {
		input, expected string
	}{
		{"hello", "hello"},
		{"world", "world"},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.input, func(t *testing.T) {
			if tt.input != tt.expected {
				t.Errorf("[want] %v\t[got] %v", tt.expected, tt.input)
			}
		})
	}
}

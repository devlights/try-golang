package lib

import "testing"

func TestFn(t *testing.T) {
	testcases := []struct {
		name string
		in   int
		out  NumberOfDigits
	}{
		{"1", 1, NumberOfDigits(1)},
		{"10", 10, NumberOfDigits(2)},
		{"100", 100, NumberOfDigits(3)},
		{"9999", 9999, NumberOfDigits(4)},
		{"123456789", 123456789, NumberOfDigits(9)},
		{"1234567890", 1234567890, NumberOfDigits(10)},
	}

	for _, tc := range testcases {
		tc := tc

		if v := UseToString(tc.in); v != tc.out {
			t.Errorf("%s-UseToString: [want] %v\t[got] %v", tc.name, tc.out, v)
		}

		if v := UseLog10(tc.in); v != tc.out {
			t.Errorf("%s-UseLog10: [want] %v\t[got] %v", tc.name, tc.out, v)
		}
	}
}

func BenchmarkUseToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseToString(123456789)
	}
}

func BenchmarkUseLog10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseLog10(123456789)
	}
}

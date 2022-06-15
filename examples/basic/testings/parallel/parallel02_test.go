package parallel

import (
	"testing"
	"time"
)

func Test_Parallel_02(t *testing.T) {
	t.Log("setup")

	tests := []struct {
		input, expected string
		delay           time.Duration
	}{
		{"hello", "hello", 800 * time.Millisecond},
		{"world", "world", 300 * time.Millisecond},
	}

	// teardownが正しい順序で呼ばれるようにするために
	// テストの実行全体を t.Run() で囲む
	t.Run("testcase", func(t *testing.T) {
		for _, tt := range tests {
			tt := tt
			t.Run(tt.input, func(t *testing.T) {
				t.Parallel()

				t.Logf("\tcase: [%s] setup", t.Name())
				defer func() { t.Logf("\tcase: [%s] teardown", t.Name()) }()

				time.Sleep(tt.delay)

				if tt.input != tt.expected {
					t.Errorf("[want] %v\t[got] %v", tt.expected, tt.input)
				}
			})
		}
	})

	t.Log("teardown")
}

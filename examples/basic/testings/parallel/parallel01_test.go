package parallel

import "testing"

func Test_Parallel_01(t *testing.T) {
	t.Log("setup")
	
	tests := []struct{
		input, expected string
	}{
		{"hello", "hello"},
		{"world", "world"},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.input, func(t *testing.T){
			t.Parallel()
			if tt.input != tt.expected {
				t.Errorf("[want] %v\t[got] %v", tt.expected, tt.input)
			}
		})
	}

	// このサンプルではテストを並行実行しているが
	// teardownの呼び方が良くない。これでは平行実行されている最中に
	// 以下のコードブロックが呼ばれてしまう。
	// 並行実行しながら、ちゃんと setup-teardownをするには parallel02_test.go のようにする
	t.Log("teardown")
}
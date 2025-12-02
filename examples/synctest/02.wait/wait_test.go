package waittest

import (
	"sync/atomic"
	"testing"
	"testing/synctest"
)

func getValueChan() <-chan int {
	var (
		out = make(chan int, 1)
	)
	go func(ch chan<- int) {
		defer close(ch)
		ch <- 999
	}(out)

	return out
}

var value atomic.Int32

func updateValue() {
	go func() {
		value.Store(999)
	}()
}

func getValue() int {
	return int(value.Load())
}

func genValues(count int) <-chan int {
	out := make(chan int)

	go func(ch chan<- int, count int) {
		defer close(ch)
		for i := range count {
			ch <- i
		}
	}(out, count)

	return out
}

// TestValueChan は、非同期処理を行う関数からチャネルを受け取り値が正しいかテストします。
// このテストでは、非同期処理であるが、チャネルを返す関数となっているため、関数内部で発行されたゴルーチンの
// 終了を確実に待機します。なので、テストコード側は当該関数内部の並行処理について何も知る必要がありません。
func TestValueChan(t *testing.T) {
	var (
		want = 999
		got  = <-getValueChan() // ここで値が返ってきた時点で確実にゴルーチンの処理が実行されている
	)
	if want != got {
		t.Errorf("[want] %v\t[got] %v", want, got)
	}
}

// TestNaiveValue は、ナイーブな値（atomic.Int32として宣言されたvalue）を更新し、更新後の結果が一致するかテストします。
// チャネルを使った場合と異なり、この値は単純に atomic.Int32 として宣言されているだけのため更新及び取得に関してのアトミック操作は
// 保証しますが、非同期処理の順序は保証しません。このテストでは内部で非同期処理の待ち合わせを全く行っていないため
// 望みの値と合致するかどうかは完全に「可能性」の問題になります。（そして、ほぼ合致しません）
func TestNaiveValueFail(t *testing.T) {
	updateValue()

	var (
		want = 999
		got  = getValue() // このタイミングで取得した値が更新後なのかどうかは不確定な状態
	)
	if want != got {
		t.Errorf("[want] %v\t[got] %v", want, got)
	}
}

// TestNaiveValueWithSynctest は、TestNaiveValueにて不確定な状態となる「更新用の非同期処理が実施されたかどうか」を
// 確定状態でテストするためのテストコードです。synctest.Wait()を用いることによりバブル内のすべてのゴルーチンが終了するまで
// 呼び元のゴルーチンをブロックさせることが出来るようになります。つまり、synctest.Wait()より前に実行もしくは実行予定となっている
// ゴルーチンは確定で終了した後で、後続の処理をテスト出来るようになります。
func TestNaiveValueWithSynctest(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		updateValue()

		// この呼び出しにより、バブル内の他のゴルーチンが終了するまで待機できる
		synctest.Wait()

		var (
			want = 999
			got  = getValue() // synctest.Wait()により確実に updateValue() 内のゴルーチンが終了後に実行できる
		)
		if want != got {
			t.Errorf("[want] %v\t[got] %v", want, got)
		}
	})
}

// TestCloseChanFail は、非同期処理の最後にチャネルを閉じる処理がある場合で呼び元の待機が漏れていると
// 閉じる前（defer close(ch))よりも先に後続の処理が進んでしまう場合があることを確認するテストです。
func TestCloseChanFail(t *testing.T) {
	ch := genValues(2)
	<-ch
	<-ch

	select {
	case _, ok := <-ch: // このタイミングで defer close(ch) がまだ実行されていない可能性がある (つまり ok=true の可能性)
		if ok {
			t.Errorf("channel is not closed")
		} else {
			t.Log("channel is closed")
		}
	default:
		t.Errorf("[want] channel was closed\t[got] channel is not close")
	}
}

// TestCloseChanFail は、TestCloseChanFailにsynctest.Wait()にてバブル内の
// ゴルーチン待機後に後続のテストが処理されることを確認します。
func TestCloseChanWithSynctest(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		ch := genValues(2)
		<-ch
		<-ch

		// この呼び出しにより、バブル内の他のゴルーチンが終了するまで待機できる
		synctest.Wait()

		select {
		case _, ok := <-ch: // 確定でgenValues()内のゴルーチンが完了後になるためチャネルは閉じている
			if ok {
				t.Errorf("channel is not closed")
			} else {
				t.Log("channel is closed")
			}
		default:
			t.Errorf("[want] channel was closed\t[got] channel is not close")
		}
	})
}

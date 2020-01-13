package log_

import (
	"github.com/getsentry/sentry-go"
	"sync"
	"time"
)

// SentryGoroutineBad は、Goroutineの中でSentryを使う場合に「してはいけないパターン」を表しているサンプルです。
// このサンプルのように、Hubを利用せずに直接スコープを定義して処理してはいけない。
func SentryGoroutineBad() error {
	// ----------------------------------------------------------------
	// sentry-go における goroutine 内での利用方法について
	//
	// Goroutine内で利用する場合のやり方について以下のページにコード付きで
	// 記載されている。
	//   https://docs.sentry.io/platforms/go/goroutines/
	// Goroutine内で、スコープを構成する場合は必ず *sentry.Hub を取得して
	// Hubに対して、スコープを構成、および、データのキャプチャを行うようにする。
	//
	// 以下は、わざとHubを利用せずにスコープを構成してメッセージをキャプチャ
	// しているサンプルである。実行すると、メッセージ自体はちゃんと届くが
	// 設定されているタグの値が高確率でおかしくなる。
	//
	// REFERENCES::
	//   https://docs.sentry.io/platforms/go/goroutines/
	//   https://docs.sentry.io/enriching-error-data/scopes/?platform=go
	// ----------------------------------------------------------------
	err := sentry.Init(sentry.ClientOptions{})
	if err != nil {
		return err
	}

	defer sentry.Flush(5 * time.Second)

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()

		// Hubを使わずに直接Goroutine内でスコープを編成してはいけない
		// 他のGoroutineで同じようにスコープを編成していたりするとデータ競合が
		// 発生して設定が上書きされてしまう可能性がある.
		sentry.ConfigureScope(func(s *sentry.Scope) {
			s.SetTag("sentry-goroutine-example-bad", "go#1")
		})

		for i := 0; i < 3; i++ {
			sentry.CaptureMessage("sentry.Hubを使わずにメッセージ送信 from Goroutine#1")
		}
	}()

	go func() {
		defer wg.Done()

		sentry.ConfigureScope(func(s *sentry.Scope) {
			s.SetTag("sentry-goroutine-example-bad", "go#2")
		})

		for i := 0; i < 3; i++ {
			sentry.CaptureMessage("sentry.Hubを使わずにメッセージ送信 from Goroutine#2")
		}
	}()

	wg.Wait()

	// -------------------------------------------
	// このサンプルを実行してSentryに届いた
	// 情報を確認すると、高確率で 2つ目の Goroutine の
	// sentry-goroutine-example タグの値が go#1 となる
	// 本来であれば、 go#2 とならないといけないが
	// Hubを利用せずにスコープを構成しているため
	// データが上書きされてしまっている。
	// -------------------------------------------

	return nil
}

package logging

import (
	"sync"
	"time"

	"github.com/getsentry/sentry-go"
)

// SentryGoroutineGood は、Goroutineの中でSentryを使う場合に「こうするべきパターン」を表しているサンプルです。
// このサンプルのように、Hubを利用してスコープを構成しないといけない。
func SentryGoroutineGood() error {
	// ----------------------------------------------------------------
	// sentry-go における goroutine 内での利用方法について
	//
	// Goroutine内で利用する場合のやり方について以下のページにコード付きで
	// 記載されている。
	//   https://docs.sentry.io/platforms/go/goroutines/
	// Goroutine内で、スコープを構成する場合は必ず *sentry.Hub を取得して
	// Hubに対して、スコープを構成、および、データのキャプチャを行うようにする。
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

	// 以下の２つのGoroutineは、それぞれ異なるタイミングでHubを取得しているが
	// どちらも正しい方法となっているため、好きな方を使えば良い。
	// (https://docs.sentry.io/platforms/go/goroutines/ 参照)
	//
	// 大事なのは、Goroutine内では必ずHub経由でSentryにアクセスすること。
	hub := sentry.CurrentHub().Clone()
	go func(h *sentry.Hub) {
		defer wg.Done()

		h.ConfigureScope(func(s *sentry.Scope) {
			s.SetTag("sentry-goroutine-example-good", "go#1")
		})

		for i := 0; i < 3; i++ {
			h.CaptureMessage("sentry.Hubを使ってメッセージ送信 from Goroutine#1")
		}
	}(hub)

	go func() {
		wg.Done()

		h := sentry.CurrentHub().Clone()
		h.ConfigureScope(func(s *sentry.Scope) {
			s.SetTag("sentry-goroutine-example-good", "go#2")
		})

		for i := 0; i < 3; i++ {
			h.CaptureMessage("sentry.Hubを使ってメッセージ送信 from Goroutine#2")
		}
	}()

	wg.Wait()

	// -------------------------------------------
	// このサンプルを実行してSentryに届いた
	// 情報を確認すると、正しくそれぞれのGorouineごとに
	// sentry-goroutine-example2 タグの値が設定されている
	// -------------------------------------------

	return nil
}

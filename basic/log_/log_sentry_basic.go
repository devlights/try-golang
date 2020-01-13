package log_

import (
	"github.com/getsentry/sentry-go"
	"log"
	"os"
	"time"
)

// SentryBasic は、SentryBasic のサービスをつかってエラーログを出力するサンプルです。
func SentryBasic() error {
	// ----------------------------------------------------------------
	// sentry-go について
	//
	// sentry-go パッケージは、SentryBasic (https://sentry.io/) の機能を
	// Goから利用するための公式SDK. 以前は raven-go というパッケージで
	// 提供されていたが、最近 sentry-go というパッケージで新たに公開されている模様。
	//
	// 	 https://github.com/getsentry/sentry-go
	//
	// Goのバージョンが 1.11 以降はこっちを利用することが推奨されている。
	//
	// 使い方は、どの言語でも大体同じで最初に Init 系の関数を呼んで
	// 初期設定を済ませたら、後はSentryに送りたいときに然るべき関数を
	// 呼び出すと自動的にSentryにエントリが送られる.
	//
	// 最も大事な情報である DSN については、環境変数で
	//   SENTRY_DSN
	// という名前で定義しておくと、自動で認識してくれる。
	// ----------------------------------------------------------------
	// 環境変数に「SENTRY_DSN」を定義しているので
	// ClientOptions.Dsn に明示的に設定しなくても
	// sentry-go 側で認識して取得してくれる.
	errSentryInit := sentry.Init(sentry.ClientOptions{})
	if errSentryInit != nil {
		log.Fatal(errSentryInit)
	}

	origFlags := log.Flags()
	log.SetFlags(0)
	defer log.SetFlags(origFlags)

	log.Println("sentryの初期設定完了")

	// 処理の終わりに Flush の呼び出しを入れておいて、メッセージの送信がちゃんと実行されるようにしておく
	defer sentry.Flush(5 * time.Second)

	// メッセージを送信
	sentry.CaptureMessage("SentryBasic テストメッセージ")
	log.Println("テストメッセージを送信")

	// わざと存在しないファイルを指定してエラーを発生させ
	// その情報をSentryに送る
	_, err := os.Open("thisisnotexists.txt")
	if err != nil {
		sentry.CaptureException(err)
		log.Println("エラー情報を送信")
	}

	return nil
}

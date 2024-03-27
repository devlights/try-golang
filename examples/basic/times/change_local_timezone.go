package times

import (
	"log"
	"os"
	"time"
)

// ChangeLocalTimezone -- time.Localを変更して強制的にローカルタイムゾーンを一時的に変更するサンプルです.
//
// REFERENCES
//   - https://pkg.go.dev/time#Location
//   - https://tutuz-tech.hatenablog.com/entry/2021/01/30/192956
//   - https://qiita.com/immrshc/items/a080975c6c7e23498944
func ChangeLocalTimezone() error {
	// Gitpod や Github Codespace などで time.Now() すると
	// タイムゾーンが JST ではなく UTC となる。(海外のサーバ上なので当然)
	//
	// 基本はこれで良いが、logパッケージのサンプルなどを動かすと
	// 時刻がズレすぎていて変なので、一時的にタイムゾーンを変更したい。
	//
	// timeパッケージには グローバル な Local フィールドが存在するので
	// ここを変更すれば強制的にタイムゾーンを変更できる。
	//
	// ただし、いい方法とは言えないので、実務レベルではやらない方が良いと思います。
	// 普通に環境変数 TZ に Asia/Tokyo を設定しておいてプログラムを動作させる方が良いと思います。

	var (
		l = log.New(os.Stderr, "", log.LstdFlags)
	)

	// Gitpod や Github Codespace などで実行すると以下は JST ではなく UTC となる
	l.Printf("[time.Now() -- Default] %v\n", time.Now())

	// 一時的な変更に留めたい場合は後から復活できるようにしておく
	original := time.Local
	defer func() {
		time.Local = original
	}()

	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return err
	}
	time.Local = jst

	// 切り替え後だと、ちゃんとタイムゾーンが適用される
	l.Printf("[time.Now() -- After  ] %v\n", time.Now())

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: time_changelocaltimezone

	   [Name] "time_changelocaltimezone"
	   2024/03/27 06:01:32 [time.Now() -- Default] 2024-03-27 06:01:32.531088275 +0000 UTC m=+0.730831549
	   2024/03/27 15:01:32 [time.Now() -- After  ] 2024-03-27 15:01:32.531173425 +0900 JST m=+0.730916689


	   [Elapsed] 96.63µs
	*/

}

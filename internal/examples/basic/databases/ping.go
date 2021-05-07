package databases

import (
	"context"
	"database/sql"
	"time"

	"github.com/devlights/gomy/output"
	"github.com/devlights/gomy/times"
	_ "modernc.org/sqlite"
)

// Ping は、 (*sql.DB).Ping の使い方についてのサンプルです.
//
// REFERENCES:
//   - https://pkg.go.dev/database/sql
//   - https://github.com/golang/go/wiki/SQLDrivers
//   - https://github.com/golang/go/wiki/SQLInterface
//   - http://go-database-sql.org/
//   - https://sourjp.github.io/posts/go-db/
func Ping() error {
	var (
		rootCtx          = context.Background()
		mainCtx, mainCxl = context.WithCancel(rootCtx)
	)
	defer mainCxl()

	var (
		db  *sql.DB
		err error
	)

	if db, err = sql.Open(Driver, Dsn); err != nil {
		return err
	}
	defer func() {
		if err = db.Close(); err != nil {
			output.Stderrf("[Error]", "db.Close: %s", err)
		}
	}()

	//
	// Ping -- 生存確認を行う
	//
	// 通信のPINGと同様に、対象の *sql.DB が生きているかどうかを確認する
	// まだ一度もクエリを発行していない状態で、このメソッドを呼ぶと接続が行われる
	//
	// 通信プログラムの場合と同様で、長時間接続を持続するアプリケーションを作成している場合などに
	// 接続か切れていないか確認する場合などで利用できる
	//
	// PingとPingContextの２つのメソッドがあり、前者はコンテキスト無しでブロッキングされ
	// 後者はコンテキストを渡すバージョン。
	//

	// Ping メソッド（ブロッキング）
	if err = db.Ping(); err != nil {
		return err
	}

	// PingContext メソッド（コンテキストを渡せる)
	var (
		pingCtx, pingCxl = context.WithTimeout(mainCtx, 100*time.Millisecond)
	)
	defer pingCxl()

	if err = db.PingContext(pingCtx); err != nil {
		return err
	}

	//
	// 定周期で生存確認
	//
	var (
		hbCtx, hbCxl = context.WithTimeout(mainCtx, 6*time.Second)
		interval     = 1 * time.Second
	)
	defer hbCxl()

	ctx := heartbeat(hbCtx, db, interval)

	<-hbCtx.Done()
	<-ctx.Done()

	return nil
}

func heartbeat(ctx context.Context, db *sql.DB, interval time.Duration) context.Context {
	ctx, stop := context.WithCancel(ctx)

	go func() {
		defer stop()

	LOOP:
		for {
			select {
			case <-ctx.Done():
				break LOOP
			case t := <-time.After(interval):
				pt := times.HHmmss(t)

				if err := db.Ping(); err != nil {
					output.Stderrf("[heartbeat]", "[Error]:%s (%s)\n", err, pt)
					continue
				}

				output.Stdoutf("[heartbeat]", "[Ok]: %s\n", pt)
			}
		}

		output.Stdoutl("[heartbeat]", "done")
	}()

	return ctx
}

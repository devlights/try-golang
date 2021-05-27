package databases

import (
	"database/sql"

	"github.com/devlights/gomy/output"
	// Goの標準ライブラリにはDBのドライバは含まれていないため
	// 対象となるDBのドライバをインポートしておく必要がある
	_ "modernc.org/sqlite"
)

// Open は、sql.Open 関数の使い方についてサンプルです.
//
// REFERENCES:
//   - https://pkg.go.dev/database/sql
//   - https://github.com/golang/go/wiki/SQLDrivers
//   - https://github.com/golang/go/wiki/SQLInterface
//   - http://go-database-sql.org/
//   - https://sourjp.github.io/posts/go-db/
func Open() error {
	var (
		db  *sql.DB // データベース
		err error   // エラー
	)

	//
	// ドライバとDSNを指定してデータベースとの接続
	//
	// sql.Open では、実際に接続されていない場合がある。
	// （遅延評価となっており、最初に操作を行った際に実際に接続される）
	// なので、クエリを発行する前に接続がちゃんと出来ているかどうかを
	// 確認したい場合は、 (*sql.DB).Ping or PingContext を利用して
	// 確認することも出来る。
	//
	// ファイルなどと同様に使い終わったらCloseする
	//
	if db, err = sql.Open(Driver, Dsn); err != nil {
		return err
	}

	defer func() {
		if err = db.Close(); err != nil {
			output.Stderrf("[Error]", "db.Close: %s", err)
		}
	}()

	output.Stdoutf("[sql.Open]", "%#v\n", db)

	return nil
}

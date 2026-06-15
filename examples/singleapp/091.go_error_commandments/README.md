# Go エラーハンドリング 10箇条

元URL： https://preslav.me/2026/05/19/10-golang-error-handling-commandments/

***

## 1. 汝、エラーを決して無視すべからず

### 悪い例

```go
// Go 1.26 想定
package bad

import "os"

func writeConfig(path string, data []byte) {
	// エラーを無視している（バグの温床）
	_ = os.WriteFile(path, data, 0o644)
}
```

### 良い例

```go
package good

import (
	"fmt"
	"os"
)

// 設定ファイルを書き出す関数
func writeConfig(path string, data []byte) error {
	// エラーを無視せず、そのまま受け取る
	if err := os.WriteFile(path, data, 0o644); err != nil {
		// 何をしていたかが分かるように文脈を付けて返す
		return fmt.Errorf("writing config file %s: %w", path, err)
	}
	return nil
}
```

ポイント:  
- `_ = err` は「本当に無視してよい特殊ケース」以外では使わない。  
- エラーを返せる関数では、呼び出し元に判断を委ねるのが基本。  

 [bytesizego](https://www.bytesizego.com/blog/error-handling-golang)

***

## 2. 汝、パッケージ境界においてはエラーを包みて返すべし

### 悪い例（ラップせず丸投げ）

```go
package user

import "database/sql"

type Repository struct {
	DB *sql.DB
}

func (r *Repository) FindByID(id int64) (*User, error) {
	row := r.DB.QueryRow("SELECT id, name FROM users WHERE id = $1", id)

	u := &User{}
	// どこで何をしていて失敗したのか文脈がない
	if err := row.Scan(&u.ID, &u.Name); err != nil {
		return nil, err
	}
	return u, nil
}
```

### 良い例（境界で文脈を付与）

```go
package user

import (
	"database/sql"
	"fmt"
)

type Repository struct {
	DB *sql.DB
}

func (r *Repository) FindByID(id int64) (*User, error) {
	row := r.DB.QueryRow("SELECT id, name FROM users WHERE id = $1", id)

	u := &User{}
	if err := row.Scan(&u.ID, &u.Name); err != nil {
		// 「どのリポジトリで」「どのIDを探しているときに」失敗したかが分かる
		return nil, fmt.Errorf("finding user by id %d: %w", id, err)
	}
	return u, nil
}
```

ポイント:  
- 他パッケージに出ていく直前で `fmt.Errorf("doing X: %w", err)` を挟む。  
- ログを見ると「user.Repository が id=42 の検索中に失敗」と一発で分かる。  

 [preslav](https://preslav.me/2024/06/06/error-flows-in-golang/)

***

## 3. 汝、同一パッケージ内においてはエラーを飾らずそのまま返すべし

### 悪い例（内部でラップしすぎ）

```go
package svc

import (
	"fmt"
)

func step1() error {
	// …
	return fmt.Errorf("step1 failed: %w", doSomething())
}

func step2() error {
	// …
	if err := step1(); err != nil {
		return fmt.Errorf("step2 failed: %w", err)
	}
	return nil
}

func Run() error {
	if err := step2(); err != nil {
		// ここでさらにラップ
		return fmt.Errorf("run failed: %w", err)
	}
	return nil
}
```

### 良い例（境界だけでラップ）

```go
package svc

import "fmt"

func step1() error {
	// パッケージ内の細かい関数は、そのままエラーを返すだけにする
	return doSomething()
}

func step2() error {
	// ここも内部なのでそのまま返す
	return step1()
}

func Run() error {
	if err := step2(); err != nil {
		// このパッケージの「外」に出るエラーだけ文脈を付ける
		return fmt.Errorf("running service: %w", err)
	}
	return nil
}
```

ポイント:  
- 「どこまでがパッケージ内部の詳細で、どこからが外向きの API か」を線引きする。  
- 外向きの関数だけがラップする、というルールで揃えるとノイズが減る。  

 [preslav](https://preslav.me/2024/06/06/error-flows-in-golang/)

***

## 4. 汝、エラーに物語らしむべし

### 悪い例（何が壊れたかだけを書く）

```go
return fmt.Errorf("db connection failed: %w", err)
```

### 良い例（何をしていたかを書く）

```go
// 「ユーザプロフィールをロードしていた」という行為を明示する
return fmt.Errorf("loading user profile (id=%d): %w", userID, err)
```

もう少し長い例:

```go
func LoadUserProfile(userID int64) (*UserProfile, error) {
	p, err := loadFromDB(userID)
	if err != nil {
		// 「DBが壊れた」ではなく「プロフィールを読み込もうとしていた」ことを先に書く
		return nil, fmt.Errorf("loading user profile (id=%d): %w", userID, err)
	}
	return p, nil
}
```

ポイント:  
- エラーログを読む人は「エラーが起きたこと」自体は知っている。  
- 「どんなストーリーの中で、どの行為のときに失敗したか」をメッセージにする。  

 [reddit](https://www.reddit.com/r/golang/comments/1pdzpbh/what_are_the_best_practices_for_error_handling_in/)

***

## 5. 汝、内なるエラーが既に語りし事柄を繰り返すべからず

### 悪い例（同じ情報の重ね書き）

```go
// 下層が既に SQL の詳細を持っているのに、テーブル名やカラム名を二重に書いている
return fmt.Errorf("users table unique email constraint violation for email=%s: %w", email, err)
```

### 良い例（レイヤごとに情報の粒度を分ける）

```go
// DB ドライバが constraint 名やカラム名を含んだエラーを返してくれていると仮定
func createUser(email string) error {
	if err := insertUser(email); err != nil {
		// ここでは「ユーザ登録処理で失敗した」というビジネス寄りの情報に留める
		return fmt.Errorf("creating user (email=%s): %w", email, err)
	}
	return nil
}
```

ポイント:  
- 下層: 実装寄り（SQL ステートメント、パス、システムコールなど）。  
- 上層: ビジネス寄り（ユーザ登録、在庫引き当てなど）。  
- 同じ情報を何度も書かず、それぞれのレイヤに合う情報だけを付ける。  

 [preslav](https://preslav.me/2024/06/06/error-flows-in-golang/)

***

## 6. 汝、エラーメッセージの文字列に契約を築くべからず

### 悪い例（文字列で分岐）

```go
package bad

import (
	"strings"
)

func handleErr(err error) {
	if err == nil {
		return
	}

	// メッセージの変更で簡単に壊れる
	if strings.Contains(err.Error(), "deadline exceeded") {
		// タイムアウト扱い
	} else if strings.Contains(err.Error(), "not found") {
		// NotFound扱い
	}
}
```

### 良い例（sentinel / 型＋`errors.Is` / `errors.As`）

```go
package good

import (
	"errors"
	"fmt"
)

var (
	// パッケージ内で定義する sentinel
	ErrNotFound = errors.New("resource not found")
)

// ラップするときは %w を使う
func findSomething(id int64) error {
	if id == 0 {
		// 直接 sentinel を返す
		return ErrNotFound
	}
	return nil
}

func handleErr(err error) {
	if err == nil {
		return
	}

	// sentinel で分岐
	if errors.Is(err, ErrNotFound) {
		fmt.Println("404 に相当する処理をする")
		return
	}

	fmt.Println("その他のエラーとして扱う")
}
```

ポイント:  
- `errors.Is` / `errors.As` は Go 1.13 以降の標準的なパターン。  
- ライブラリを書く側は、利用者が `errors.Is` しやすいように sentinel や型付きエラーを提供する。  

 [go](https://go.dev/blog/go1.13-errors)

***

## 7. 汝、`%w` を API の約定と心すべし

### 良い例（ラップして外に晒す）

```go
package repo

import (
	"database/sql"
	"errors"
	"fmt"
)

var ErrUserNotFound = errors.New("user not found")

type Repository struct {
	DB *sql.DB
}

func (r *Repository) FindByID(id int64) (*User, error) {
	row := r.DB.QueryRow("SELECT id, name FROM users WHERE id = $1", id)

	u := &User{}
	if err := row.Scan(&u.ID, &u.Name); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// 独自の sentinel に変換
			return nil, ErrUserNotFound
		}
		// ここでは DB のエラーをラップして公開することを許容している
		return nil, fmt.Errorf("querying user by id %d: %w", id, err)
	}
	return u, nil
}
```

### 実装詳細を晒したくない場合

```go
func (r *Repository) FindByID(id int64) (*User, error) {
	// ...
	if err := row.Scan(&u.ID, &u.Name); err != nil {
		// 外からは sql.ErrNoRows 等を見せたくない方針なら %w を使わない
		return nil, fmt.Errorf("querying user by id %d: %v", id, err) // ここではチェーンを止める
	}
	return u, nil
}
```

ポイント:  
- `%w` を使う → 「このパッケージは内側のエラーを API として公開する」ことを意味する。  
- 将来の互換性を考えるなら、どこまで `%w` で晒すかを慎重に決める。  

 [go](https://go.dev/blog/go1.13-errors)

***

## 8. 汝、外なる世界のエラーを汝自身の語彙へと訳すべし

### 悪い例（HTTP クライアントの詳細を上位に漏らす）

```go
func FetchUserFromAPI(id string) (*User, error) {
	resp, err := http.Get("https://example.com/users/" + id)
	if err != nil {
		// 上位は http パッケージの詳細を知らないと扱えない
		return nil, err
	}
	// ...
}
```

### 良い例（自分のドメインのエラーに翻訳）

```go
package api

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	ErrUserNotFound  = errors.New("user not found")
	ErrUpstreamError = errors.New("upstream API error")
)

func FetchUserFromAPI(id string) (*User, error) {
	resp, err := http.Get("https://example.com/users/" + id)
	if err != nil {
		// 通信失敗などは上流 API のエラーとしてまとめる
		return nil, fmt.Errorf("calling upstream user API: %w", err)
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		// 正常処理
	case http.StatusNotFound:
		// 404 は自分のドメインの ErrUserNotFound に変換
		return nil, ErrUserNotFound
	default:
		// それ以外は upstream の障害として扱う
		return nil, fmt.Errorf("%w: status=%d", ErrUpstreamError, resp.StatusCode)
	}

	return nil, nil
}
```

ポイント:  
- 上位レイヤーは「ユーザがいない」「外部システム障害」など、ビジネス的に意味のあるエラーだけを知ればよい。  
- 外部ライブラリの細かいエラーは adapter 層で吸収する。  

 [bytesizego](https://www.bytesizego.com/blog/error-handling-golang)

***

## 9. 汝、エラーをログしつつ同時に返すべからず

### 悪い例（二重ログ）

```go
func doSomething() error {
	if err := step(); err != nil {
		// 中間層でログ
		log.Printf("step failed: %v", err)
		return err
	}
	return nil
}

func main() {
	if err := doSomething(); err != nil {
		// 最上位でも同じエラーをログ
		log.Fatalf("fatal: %v", err)
	}
}
```

### 良い例（ログするのは最上位だけ）

```go
func doSomething() error {
	if err := step(); err != nil {
		// 中間層はログしない。文脈を付けて返すだけ
		return fmt.Errorf("running step: %w", err)
	}
	return nil
}

func main() {
	if err := doSomething(); err != nil {
		// 「このアプリケーション全体としてどう扱うか」を決める層だけがログを出す
		log.Fatalf("application failed: %v", err)
	}
}
```

ポイント:  
- 原則「log するか、返すかのどちらかひとつ」。  
- ログを集約するレイヤー（HTTP ミドルウェアや main）を決めておくと運用が楽。  

 [reddit](https://www.reddit.com/r/golang/comments/1tjmns3/the_10_go_error_handling_commandments/)

***

## 10. 汝、ゴルーチンの内にて生じたエラーを聞かれざるまま放置すべからず

### 悪い例（ゴルーチン内で握りつぶし）

```go
func startWorker() {
	go func() {
		if err := doWork(); err != nil {
			// 何もせず無視。誰もこのエラーを知ることができない
		}
	}()
}
```

### 良い例（チャネルで伝える）

```go
package worker

import (
	"context"
	"fmt"
)

func startWorker(ctx context.Context) <-chan error {
	errCh := make(chan error, 1)

	go func() {
		// ゴルーチンが終わったらチャネルを閉じる
		defer close(errCh)

		if err := doWork(ctx); err != nil {
			// 呼び出し元にエラーを伝える
			errCh <- fmt.Errorf("worker: %w", err)
			return
		}
	}()

	return errCh
}
```

呼び出し側:

```go
func main() {
	ctx := context.Background()
	errCh := startWorker(ctx)

	// どこかで待ち受けて、エラーを観測する
	if err := <-errCh; err != nil {
		log.Fatalf("worker failed: %v", err)
	}
}
```

### `errgroup` を使った良い例

```go
import (
	"context"
	"fmt"

	"golang.org/x/sync/errgroup"
)

func run(ctx context.Context) error {
	g, ctx := errgroup.WithContext(ctx)

	// 1つ目のゴルーチン
	g.Go(func() error {
		if err := doWork1(ctx); err != nil {
			return fmt.Errorf("work1: %w", err)
		}
		return nil
	})

	// 2つ目のゴルーチン
	g.Go(func() error {
		if err := doWork2(ctx); err != nil {
			return fmt.Errorf("work2: %w", err)
		}
		return nil
	})

	// どれかがエラーを返すと g.Wait() がそのエラーを返してくれる
	if err := g.Wait(); err != nil {
		return fmt.Errorf("concurrent run: %w", err)
	}
	return nil
}
```

ポイント:  
- ゴルーチンは戻り値を返せないので、「エラーをどこに届けるか」を先に設計する。  
- `golang.org/x/sync/errgroup` は「並列処理＋エラー伝播」をまとめて扱うのに便利。  

 [codingexplorations](https://www.codingexplorations.com/blog/mastering-concurrency-in-go-with-errgroup-simplifying-goroutine-management)

***

## （おまけ）`context.Canceled` / `context.DeadlineExceeded` を見分ける

### 良い例（最外周で特別扱い）

```go
import (
	"context"
	"errors"
	"log"
)

func main() {
	ctx := context.Background()

	if err := run(ctx); err != nil {
		// キャンセル・タイムアウトは「想定された終了」として扱う
		if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
			log.Printf("request canceled: %v", err)
			return
		}
		// それ以外は本当の障害として扱う
		log.Fatalf("server failed: %v", err)
	}
}
```

ポイント:  
- キャンセル／タイムアウトは「ユーザや上位の意思」で止めた結果であることも多い。  
- これを通常のエラーと同列に扱うと、エラーログがノイズで溢れる。  

 [gopherguides](https://www.gopherguides.com/golang-fundamentals-book/12-context/errors)


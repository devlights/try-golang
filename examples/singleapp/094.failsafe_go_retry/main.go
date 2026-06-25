// failsafe-goライブラリのサンプルです。
//
// failsafe-go は、
//
//   - リトライ
//
//     失敗した実行を、設定した回数・間隔・バックオフ戦略に基づき再試行するポリシー
//
//     ネットワークエラーや一時的な障害に対して、クライアント側で自動再試行を行う用途
//
//   - サーキットブレーカー
//
//     連続／一定割合以上の失敗を検出して「回路を開き」、一定期間は即座に失敗を返して負荷や連鎖障害を防ぐポリシー
//
//     一定時間経過後に「半開状態」で試行実行を行い、成功が閾値を満たせば再び閉じる（正常に戻る）という典型的な CB パターン
//
//   - レートリミット
//
//     一定時間あたりの許可実行数を制御し、呼び出し頻度を制限するポリシー
//
//     外部サービスや API に対して「1 秒あたり N 回」「分あたり M 回」といった制限をかける用途
//
//   - タイムアウト
//
//     実行時間が一定時間を超えた場合に、実行中の処理をタイムアウトとして失敗扱いにするポリシー
//
//     固定タイムアウト：単純に「N 秒」で打ち切る設定
//
//   - フォールバック
//
//     失敗時に 別の処理（バックアップ処理）を実行して結果を返すポリシー
//
//     Retry や Circuit Breaker と合わせて、「再試行の末に失敗したら最後にフォールバックへ」といったシナリオを構成
//
//   - ヘッジ
//
//     ヘッジリクエストと呼ばれるパターンで、先に成功した結果を採用し、遅い方はキャンセルといったレイテンシ削減テクニックを実装するポリシー
//
//     一定遅延後にヘッジリクエストを投げることで、平均レイテンシを下げつつも過剰な重複負荷を避ける構成
//
//   - バルクヘッド
//
//     船の「水密隔壁」になぞらえたパターンで、同時実行数やキュー長を制限して、特定の依存先への負荷が他の処理に波及しないようにするポリシー
//
//     同時実行数制限：指定した上限数を超えると bulkhead.ErrFull のようなエラーで新規実行を拒否
//
//   - キャッシュ
//
//     実行結果をキャッシュして再利用し、同一入力に対する再実行を省略するポリシー
//
//     キー付きキャッシュ：入力やコンテキストに基づきキーを生成し、それに紐づく結果をキャッシュ
//
// といった耐障害・レジリエンスパターンをポリシーとして合成しながら使えるライブラリ
//
// 使い方として、まずPolicyを生成し、次にExecutorを生成する。
// 生成したExecutorに実処理を渡して、ポリシーに従った処理をしてもらう形となる。
//
// # REFERENCES
//   - https://github.com/failsafe-go/failsafe-go
//   - https://failsafe-go.dev/retry/
package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/failsafe-go/failsafe-go"
	"github.com/failsafe-go/failsafe-go/retrypolicy"
)

const (
	addr = "localhost:8888" // リトライを試すのでListenしないこと
)

func main() {
	log.SetFlags(log.Ltime)

	var (
		// 失敗時に１秒インターバルで最大３回のリトライを行うポリシー
		policy   = retrypolicy.NewBuilder[net.Conn]().WithDelay(time.Second).WithMaxRetries(3).Build()
		executor = failsafe.With(policy)
		dialFn   = func() (net.Conn, error) {
			conn, err := net.Dial("tcp", "localhost:8888")
			if err != nil {
				log.Printf("dialFn: %v", err)
				return nil, fmt.Errorf("接続失敗： %s (%w)", addr, err)
			}

			return conn, nil
		}
	)

	conn, err := executor.Get(dialFn)
	if err != nil {
		log.Printf("retry-over: %[1]v (%[1]T)", err)
		return
	}
	defer conn.Close()

	log.Println("dial成功（ここには来ない）")
}

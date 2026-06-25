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
	//
	// failsafe-go は、
	//   - リトライ
	//   - サーキットブレーカー
	//   - レートリミット
	//   - タイムアウト
	//   - フォールバック
	//   - ヘッジ
	//   - バルクヘッド
	//   - キャッシュ
	// といった耐障害・レジリエンスパターンをポリシーとして合成しながら使えるライブラリ
	//
	// 使い方として、まずPolicyを生成し、次にExecutorを生成する。
	// 生成したExecutorに実処理を渡して、ポリシーに従った処理をしてもらう形となる。
	//
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

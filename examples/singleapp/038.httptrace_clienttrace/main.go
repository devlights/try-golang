package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"net/http/httptrace"
	"time"
)

var start time.Time

func init() {
	log.SetFlags(log.Lmicroseconds)
}

func trace() *httptrace.ClientTrace {
	return &httptrace.ClientTrace{
		// 名前解決を開始した時
		DNSStart: func(info httptrace.DNSStartInfo) {
			log.Printf("名前解決 開始: %s\n", info.Host)
		},
		// 名前解決が完了した時
		DNSDone: func(info httptrace.DNSDoneInfo) {
			log.Printf("名前解決 完了: %v (duration: %s)\n", info.Addrs, time.Since(start))
		},
		// TCP 接続開始時
		ConnectStart: func(network, addr string) {
			log.Printf("TCP接続 開始: %s...\n", addr)
		},
		// TCP 接続完了（または失敗）時
		ConnectDone: func(network, addr string, err error) {
			if err != nil {
				log.Printf("TCP接続 エラー: %v\n", err)
			} else {
				log.Printf("TCP接続 完了: %s\n", addr)
			}
		},
		// HTTPS の TLS ハンドシェイク開始時
		TLSHandshakeStart: func() {
			log.Printf("TLSハンドシェイク 開始\n")
		},
		// TLS ハンドシェイク完了時
		TLSHandshakeDone: func(state tls.ConnectionState, err error) {
			if err != nil {
				log.Printf("TLSハンドシェイク エラー: %s", err)
			} else {
				log.Printf("TLSハンドシェイク 完了: バージョン=%x\n", state.Version)
			}
		},
		// リクエスト書き込み完了時
		WroteRequest: func(reqInfo httptrace.WroteRequestInfo) {
			if reqInfo.Err != nil {
				log.Printf("リクエスト書き込み エラー: %v\n", reqInfo.Err)
			} else {
				log.Printf("リクエスト書き込み 完了\n")
			}
		},
		// 実際に使用されるコネクションが確定した時
		// 	- 新規接続か
		// 	- 再利用か
		GotConn: func(info httptrace.GotConnInfo) {
			if info.Reused {
				log.Printf("コネクションを再利用 (idle: %s)\n", info.IdleTime)
			} else {
				log.Printf("新規接続\n")
			}
		},
	}
}

func main() {
	var (
		ctx, cxl = context.WithTimeout(context.Background(), 5*time.Second)
		err      error
	)
	defer cxl()

	if err = run(ctx); err != nil {
		panic(err)
	}
}

func run(pCtx context.Context) error {
	start = time.Now()

	var (
		req *http.Request
		err error
	)
	if req, err = http.NewRequestWithContext(pCtx, "GET", "https://example.com", nil); err != nil {
		return fmt.Errorf("Make new http.request failed: %w", err)
	}

	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace()))

	var (
		client = new(http.Client)
		resp   *http.Response
	)
	if resp, err = client.Do(req); err != nil {
		return fmt.Errorf("Do request failed: %w", err)
	}
	defer resp.Body.Close()

	log.Printf("Response status: %s (total time: %s)\n", resp.Status, time.Since(start))

	return nil
}

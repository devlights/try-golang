package logging

import (
	"log"
	"os"
	"sync"
)

// clone は、指定されたロガーのWriter()を受け継いで新たなロガーを生成して返します.
func clone(l *log.Logger, prefix string, flags int) *log.Logger {
	return log.New(l.Writer(), prefix, flags)
}

// NewLogger は、log.New() の挙動を確認するサンプルです.
func NewLogger() error {
	// ----------------------------------------------------------------
	// log.New() について
	//
	// log.New() を使うことで標準ロガー以外に自分用のロガーを生成することが出来る
	// 後の使い方は標準ロガーと同じ.
	// ----------------------------------------------------------------
	writer := os.Stderr
	rootLogger := log.New(writer, "", 0)
	rootLogger.Println("rootLogger#1")

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()

		logger := clone(rootLogger, "[go#1] ", 0)
		logger.Println("logger#1")
	}()

	go func() {
		defer wg.Done()

		logger := clone(rootLogger, "[go#2] ", 0)
		logger.Println("logger#2")
	}()

	wg.Wait()

	defer rootLogger.Println("rootLogger#2")

	return nil
}

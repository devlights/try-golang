package logging

import (
	"log"
	"time"

	"gopkg.in/natefinch/lumberjack.v2"
)

// DateRolling は、```github.com/natefinch/lumberjack``` を用いて日付でローリングするログを実装するサンプルです.
//
// # REFERENCES
//   - https://github.com/natefinch/lumberjack
//   - https://stackoverflow.com/a/36163876
//   - https://stackoverflow.com/a/33109737
func DateRolling() error {
	log.SetFlags(0)

	//
	// log.SetOutputにlumberjackのLoggerを設定する
	// lumberjack.Loggerは、io.WriteCloser を実装している
	//
	log.SetOutput(&lumberjack.Logger{
		Filename:   "daterolling.log",
		MaxSize:    1,
		MaxBackups: 1,
		MaxAge:     1,
	})

	//
	// 後は普通に利用すれば良い
	//
	const (
		msg = "helloworld"
		mb  = 1_000_000
		cnt = mb/len(msg) + 1
	)

	for i := 0; i < 5; i++ {
		for j := 0; j < cnt; j++ {
			log.Println(msg)
		}

		time.Sleep(100 * time.Millisecond)
	}

	return nil
}

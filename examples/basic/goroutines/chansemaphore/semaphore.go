// Package chansemaphore は、チャネルでセマフォの動作を実現するサンプルが配置されています
package chansemaphore

type (
	// Semaphore は、共有する資源にアクセスするのを制御する抽象データ型です.
	Semaphore interface {
		// Acquire は、共有資源にアクセスする権利を獲得します
		Acquire()
		// Release は、獲得している権利を開放します
		Release()
	}
)

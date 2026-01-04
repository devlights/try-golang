package osop

import (
	"fmt"
	"os"
	"os/user"
)

// HomeDir は、os.UserHomeDir()のサンプルです。
// Go 1.12から追加された関数。user.Current.HomeDirを使わなくても良くなる。
//
// osパッケージ側に以下の関数が用意されている
//
// - os.UserHomeDir
// - os.UserConfigDir
// - os.UserCacheDir
//
// # REFERENCES
//   - https://pkg.go.dev/os@go1.25.5#UserHomeDir
//   - https://text.baldanders.info/golang/no-need-go-homedir/
func HomeDir() error {
	// go 1.12より前は以下の様にしていた。
	u, err := user.Current()
	if err != nil {
		return fmt.Errorf("user.Current: %w", err)
	}

	fmt.Printf("user.Current.HomeDir: %s\n", u.HomeDir)

	// go 1.12以降は以下で取得できる
	homedir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("os.UserHomeDir: %w", err)
	}

	fmt.Printf("os.HomeDir: %s\n", homedir)

	return nil
}

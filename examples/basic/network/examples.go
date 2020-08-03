package network

import (
	"github.com/devlights/try-golang/mappings"
)

type (
	register struct{}
)

// NewRegister -- このパッケージ用のサンプルを登録する mappings.Register を生成します。
func NewRegister() mappings.Register {
	return new(register)
}

// Regist -- 登録します.
func (r *register) Regist(m mappings.ExampleMapping) {
	m["ipaddress_parse"] = IpAddressParse
	m["ssh_no_privkey_passphrase"] = SSHNoPrivKeyPassPhrase
	m["ssh_with_privkey_passphrase"] = SSHWithPrivKeyPassPhrase
	m["ssh_close_after_run"] = SSHSessionCloseAfterRun
}
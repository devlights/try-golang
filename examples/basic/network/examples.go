package network

import "github.com/devlights/try-golang/mapping"

type (
	register struct{}
)

// NewRegister -- このパッケージ用のサンプルを登録する mapping.Register を生成します。
func NewRegister() mapping.Register {
	return new(register)
}

// Regist -- 登録します.
func (r *register) Regist(m mapping.ExampleMapping) {
	m["network_ipaddress_parse"] = IpAddressParse
	m["network_ssh_no_privkey_passphrase"] = SSHNoPrivKeyPassPhrase
	m["network_ssh_with_privkey_passphrase"] = SSHWithPrivKeyPassPhrase
	m["network_ssh_close_after_run"] = SSHSessionCloseAfterRun
	m["network_http_get"] = HttpGet
	m["network_join_host_port"] = JoinHostPort
}

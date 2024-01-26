# サンプルリスト

このディレクトリには以下のサンプルがあります。

| file                           | example name                        | note                                                                                       |
| ------------------------------ | ----------------------------------- | ------------------------------------------------------------------------------------------ |
| ipaddress_parse.go             | network_ipaddress_parse             | net.ParseIP() の サンプルです.                                                             |
| ssh_no_privkey_passphrase.go   | network_ssh_no_privkey_passphrase   | 秘密鍵のパスフレーズ無しのSSH接続サンプルです.                                             |
| ssh_with_privkey_passphrase.go | network_ssh_with_privkey_passphrase | 秘密鍵のパスフレーズありのSSH接続サンプルです.                                             |
| ssh_close_after_run.go         | network_ssh_close_after_run         | ssh.Run() を呼んだ後に ssh.Close() を呼ぶと io.EOF が返却されることを確認するサンプルです. |
| http_get.go                    | network_http_get                    | http.Get() の サンプルです.                                                                |
| join_host_port.go              | network_join_host_port              | JoinHostPort は、net.JoinHostPort のサンプルです                                           |
| split_host_port.go             | network_split_host_port             | SplitJoinPort は、net.SplitHostPort のサンプルです                                         |
| lookup_port.go                 | network_lookup_port                 | LookupPort は、 net.LookupPort() のサンプルです                                            |

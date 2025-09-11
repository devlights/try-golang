# TCP通信でのRST送信のサンプル

Goでは ```*net.TCPConn.SetLinger()``` にて ```SO_LINGER``` が設定出来る。
値を ```0``` で設定すると切断時にRST送信が行われる。（これはC言語でも同じ）

```go
tcpConn, ok := conn.(*net.TCPConn)
if ok {
    tcpConn.SetLinger(0)
}

conn.Close() // RST送信
```

```sh
$ task
task: [build] go build -o app main.go
task: [run-fin] sudo tcpdump -i lo -n 'tcp port 8888' -S &
task: [run-fin] sleep 1
tcpdump: verbose output suppressed, use -v[v]... for full protocol decode
listening on lo, link-type EN10MB (Ethernet), snapshot length 262144 bytes
task: [run-fin] ./app -server &
task: [run-fin] ./app
task: [run-fin] sleep 2
09:22:02.504826 IP 127.0.0.1.32782 > 127.0.0.1.8888: Flags [S], seq 1561972627, win 43690, options [mss 65495,sackOK,TS val 3209706398 ecr 0,nop,wscale 7], length 0
09:22:02.504835 IP 127.0.0.1.8888 > 127.0.0.1.32782: Flags [S.], seq 2366941560, ack 1561972628, win 43690, options [mss 65495,sackOK,TS val 3209706398 ecr 3209706398,nop,wscale 7], length 0
09:22:02.504844 IP 127.0.0.1.32782 > 127.0.0.1.8888: Flags [.], ack 2366941561, win 342, options [nop,nop,TS val 3209706398 ecr 3209706398], length 0
09:22:02.504914 IP 127.0.0.1.8888 > 127.0.0.1.32782: Flags [P.], seq 2366941561:2366941566, ack 1561972628, win 342, options [nop,nop,TS val 3209706398 ecr 3209706398], length 5
09:22:02.504924 IP 127.0.0.1.32782 > 127.0.0.1.8888: Flags [.], ack 2366941566, win 342, options [nop,nop,TS val 3209706398 ecr 3209706398], length 0
09:22:02.504944 IP 127.0.0.1.32782 > 127.0.0.1.8888: Flags [F.], seq 1561972628, ack 2366941566, win 342, options [nop,nop,TS val 3209706398 ecr 3209706398], length 0
09:22:02.504982 IP 127.0.0.1.8888 > 127.0.0.1.32782: Flags [F.], seq 2366941566, ack 1561972629, win 342, options [nop,nop,TS val 3209706398 ecr 3209706398], length 0
09:22:02.504995 IP 127.0.0.1.32782 > 127.0.0.1.8888: Flags [.], ack 2366941567, win 342, options [nop,nop,TS val 3209706398 ecr 3209706398], length 0
task: [run-fin] sudo pkill tcpdump

8 packets captured
16 packets received by filter
0 packets dropped by kernel
task: [run] sleep 2
task: [run-rst] sudo tcpdump -i lo -n 'tcp port 8888' -S &
task: [run-rst] sleep 1
tcpdump: verbose output suppressed, use -v[v]... for full protocol decode
listening on lo, link-type EN10MB (Ethernet), snapshot length 262144 bytes
task: [run-rst] ./app -server -rst &
task: [run-rst] ./app -rst
task: [run-rst] sleep 2
09:22:07.566462 IP 127.0.0.1.50180 > 127.0.0.1.8888: Flags [S], seq 3347564028, win 43690, options [mss 65495,sackOK,TS val 3209711459 ecr 0,nop,wscale 7], length 0
09:22:07.566476 IP 127.0.0.1.8888 > 127.0.0.1.50180: Flags [S.], seq 671550310, ack 3347564029, win 43690, options [mss 65495,sackOK,TS val 3209711459 ecr 3209711459,nop,wscale 7], length 0
09:22:07.566485 IP 127.0.0.1.50180 > 127.0.0.1.8888: Flags [.], ack 671550311, win 342, options [nop,nop,TS val 3209711459 ecr 3209711459], length 0
09:22:07.566556 IP 127.0.0.1.8888 > 127.0.0.1.50180: Flags [P.], seq 671550311:671550316, ack 3347564029, win 342, options [nop,nop,TS val 3209711459 ecr 3209711459], length 5
09:22:07.566566 IP 127.0.0.1.50180 > 127.0.0.1.8888: Flags [.], ack 671550316, win 342, options [nop,nop,TS val 3209711459 ecr 3209711459], length 0
09:22:07.566589 IP 127.0.0.1.50180 > 127.0.0.1.8888: Flags [R.], seq 3347564029, ack 671550316, win 342, options [nop,nop,TS val 3209711459 ecr 3209711459], length 0
task: [run-rst] sudo pkill tcpdump

6 packets captured
12 packets received by filter
0 packets dropped by kernel
```

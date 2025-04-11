# io.ReadFull()のサンプル

```sh
$ task
task: [build] go build -o app .
task: [run] ./app -timeout 100ms
03:01:14.980749 [C] recv start
03:01:15.081186 netErr.Timeout()
task: [run] ./app -timeout 3s
03:01:15.091656 [C] recv start
03:01:15.592085 io.ErrUnexpectedEOF
task: [run] ./app -length 10 -timeout 3s
03:01:15.602765 [C] recv start
03:01:15.602902 [C] data=(hhhh)
03:01:15.602912 [C] data=(hhhh)
03:01:16.103520 io.ErrUnexpectedEOF
task: [run] ./app -bufsize 6 -length 12 -timeout 3s
03:01:16.113780 [C] recv start
03:01:16.113887 [C] data=(hhhhhh)
03:01:16.113901 [C] data=(hhhhhh)
03:01:16.614526 io.EOF
```


- [io.ReadFull](https://pkg.go.dev/io@latest#ReadFull)

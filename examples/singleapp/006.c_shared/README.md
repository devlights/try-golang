# Go側からsoファイルを作成してPythonとCで利用

```go build``` 時に ```-buildmode=c-shared``` を付与することで so ファイルと ヘッダファイル が生成される。

```sh
$ task run
task: [build] go build -buildmode=c-shared -o libgoadd.so main.go
task: [show] ls -lh libgoadd.so
-rw-r--r-- 1 gitpod gitpod 2.0M Nov 14 07:21 libgoadd.so
task: [show] file libgoadd.so
libgoadd.so: ELF 64-bit LSB shared object, x86-64, version 1 (SYSV), dynamically linked, BuildID[sha1]=7c44e5b5941164345845b1bb8d3103399b4870d8, with debug_info, not stripped
task: [show] ldd libgoadd.so
        linux-vdso.so.1 (0x00007ffd2b9d6000)
        libc.so.6 => /lib/x86_64-linux-gnu/libc.so.6 (0x00007fd4f6bf7000)
        /lib64/ld-linux-x86-64.so.2 (0x00007fd4f6fa9000)
task: [show] nm -D libgoadd.so | grep 'T Add'
00000000000a1e00 T Add
task: [use] python3 use.py
[FROM GOLANG] library loaded!
[FROM PYTHON] 333
task: [use] gcc -o use-c use.c -L . -l goadd
task: [use] LD_LIBRARY_PATH=. ./use-c
[FROM GOLANG] library loaded!
[FROM C] 333
```

## 参考情報

- https://pkg.go.dev/cmd/go#hdr-Build_modes
- https://qiita.com/yanolab/items/1e0dd7fd27f19f697285
- https://linuxcommand.net/nm/
- https://moznion.hatenadiary.com/entry/2022/11/06/142618
- https://c.perlzemi.com/blog/20210628105352.html

# ビルド時のldflagsの-s -wの指定について

Go 1.22以降、`-s`フラグが`-w`を暗黙的に含むようになったため、以前のように `-ldflags="-s -w"`とする必要がなくなった。

`-s`フラグだけの指定で良い。

- `-s`: シンボルテーブルの生成を抑制
- `-w`: DWARFデバッグ情報の生成を抑制

`-w=0`とすれば打ち消し可能。

なので、`-s -w=0`とすれば、「シンボルテーブルの生成なし、DWARFデバッグ情報あり」となる。

- https://go.dev/doc/go1.22#linker

## サンプルの実行結果

```
$ task
task: [build] go build -o app0 main.go
task: [build] go build -ldflags="-s" -o app1 main.go
task: [build] go build -ldflags="-s -w=0" -o app2 main.go
task: [verify] file app0 app1 app2
app0: ELF 64-bit LSB executable, ARM aarch64, version 1 (SYSV), statically linked, ..., with debug_info, not stripped
app1: ELF 64-bit LSB executable, ARM aarch64, version 1 (SYSV), statically linked, ..., stripped
app2: ELF 64-bit LSB executable, ARM aarch64, version 1 (SYSV), statically linked, ..., with debug_info, not stripped
task: [verify] readelf -S app0 | grep -E '\.debug|\.symtab|\.strtab'
  [15] .debug_abbrev     PROGBITS         0000000000000000  00180000
  [16] .debug_line       PROGBITS         0000000000000000  00180160
  [17] .debug_frame      PROGBITS         0000000000000000  001a1ad7
  [18] .debug_gdb_s[...] PROGBITS         0000000000000000  001a8c29
  [19] .debug_info       PROGBITS         0000000000000000  001a8c59
  [20] .debug_loclists   PROGBITS         0000000000000000  001f1d60
  [21] .debug_rnglists   PROGBITS         0000000000000000  0020bdfe
  [22] .debug_addr       PROGBITS         0000000000000000  0021a4ce
  [23] .symtab           SYMTAB           0000000000000000  0021b548
  [24] .strtab           STRTAB           0000000000000000  00229a20
task: [verify] readelf -S app1 | grep -E '\.debug|\.symtab|\.strtab'
task: [verify] readelf -S app2 | grep -E '\.debug|\.symtab|\.strtab'
  [15] .debug_abbrev     PROGBITS         0000000000000000  00180000
  [16] .debug_line       PROGBITS         0000000000000000  00180160
  [17] .debug_frame      PROGBITS         0000000000000000  001a1ad7
  [18] .debug_gdb_s[...] PROGBITS         0000000000000000  001a8c29
  [19] .debug_info       PROGBITS         0000000000000000  001a8c59
  [20] .debug_loclists   PROGBITS         0000000000000000  001f1d60
  [21] .debug_rnglists   PROGBITS         0000000000000000  0020bdfe
  [22] .debug_addr       PROGBITS         0000000000000000  0021a4ce
```

app1は想定通り`stripped`と表示されている。なので、シンボルテーブル無しでDWARFも無し。
app2の方は、想定ではシンボルテーブル無しでDWARF有りとなるはずであるば、`not stripped`と表示されている。
これは、fileコマンドがDWARFが存在しているので、このような表示になっている。

readelfコマンドで各テーブルの情報を確認してみると、app2の方は想定通りsymtabとstrtabは存在せずdebugのみとなっている。


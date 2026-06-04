# これは何？

Go 1.18 で追加された ```strings.Clone()``` を利用したサンプルです。( ```bytes.Clone()``` は Go 1.20 で追加)

内部で大きな文字列を確保している状態で、それらの部分文字列を別の場所に確保する処理を実施しています。

現状（2023-12-05 現在）のGoの標準コンパイラでは、元の文字列と部分文字列は同じメモリデータを共有するので

部分文字列をシャローコピーして別のストアに保持したままだと、メモリが開放されません。

```strings.Clone()``` を利用することにより、ディープコピーが行われるので、メモリが開放されるようになります。

## 実行例

```sh
$ task
task: [build] go build -o app main.go
task: [run-not-use-clone] ./app
Title           HeapAlloc       HeapObjects
[start     ]      192792             144
[gen       ]    11482528            4576
[store     ]    11487008            4588
[checkpoint]     8471296            1363
[checkpoint]     8475728            1372
[checkpoint]     8475728            1372
[checkpoint]     8475728            1372
[checkpoint]     8475736            1373
task: [run-use-clone] ./app -use
Title           HeapAlloc       HeapObjects
[start     ]      192824             144
[gen       ]    11497632            4607
[store     ]    11507440            4952
[checkpoint]      296112             724
[checkpoint]      300536             732
[checkpoint]      300536             732
[checkpoint]      300544             733
[checkpoint]      300544             733
```

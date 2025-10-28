# これは何？

このサンプルプログラムは、cgoを利用してGo言語とC言語の間でデータを連携させる方法、特にGo 1.17で導入された`unsafe.Slice`とGo 1.20で導入された`unsafe.SliceData`を活用した効率的なメモリアクセス方法を具体的に示すものです。

Cの関数からGoの関数を呼び出し、Go側で受け取ったデータを処理して、再びCの関数に処理結果を返す、という一連の流れを実装しています。

## 処理の流れ

このプログラムは、以下の順序で処理が実行されます。

1.  **`main.go:main()`**
    *   プログラムのエントリポイント。C言語側で定義された`c_func()`を呼び出します。

2.  **`c.go:c_func()`**
    *   スタック上に文字列 `"helloworld"` を確保します。
    *   Go側でエクスポートされている`go_func()`を、文字列のポインタとサイズを引数にして呼び出します。

3.  **`export.go:go_func()`**
    *   Cから渡されたポインタ (`*C.char`) とサイズ (`C.size_t`) を `unsafe.Slice` を使ってGoのスライス (`[]byte`) に変換します。この操作はメモリコピーを発生させず、Cのメモリ領域を直接参照します。
    *   安全のため、Cのメモリを直接変更するのではなく、Goの管理するメモリにデータをコピーします。
    *   コピーしたデータに対して文字列の反転処理を行います。
    *   処理後のGoスライスを `unsafe.SliceData` を使ってCで扱えるポインタ形式に変換し、C側の`c_func2()`を呼び出します。

4.  **`c.go:c_func2()`**
    *   Goから渡されたデータを受け取り、標準出力に表示します。

## 技術詳細

### CからGoへのデータ受け渡し: `unsafe.Slice`

Goの関数 (`go_func`) がCからデータを受け取る際、`unsafe.Slice` を利用してパフォーマンスを向上させています。

```go
//export go_func
func go_func(s *C.char, n C.size_t) {
	var (
		sPtr   = unsafe.Pointer(s)
		cSlice = unsafe.Slice((*byte)(sPtr), n) // cSliceはC側のスタック変数を指している
	)
	fmt.Printf("[Go] %s", cSlice)
    // ...
}
```

- `*C.char` を `unsafe.Pointer` を経由してGoの `*byte` に変換します。
- `unsafe.Slice` は、このポインタとデータ長 `n` を基に、Goのスライスヘッダを生成します。
- この結果得られる `cSlice` は、C言語側のメモリ領域を直接指し示すスライスとなり、**余計なメモリコピーが発生しません（ゼロコピー）**。

**【重要】注意点:**
`cSlice`が参照しているのはCのスタックメモリです。Goのガベージコレクタの管理外であり、関数を抜けると無効になる可能性があります。Go側でこのデータを永続化したり変更したりする場合は、必ずGoが管理するメモリに`copy()`で複製してから操作する必要があります。

### GoからCへのデータ受け渡し: `unsafe.SliceData`

Goで処理したデータをCの関数に渡す際には、`unsafe.SliceData` を利用します。

```go
// ...
	// C側の関数に渡すための準備
	var (
		bytePtr = unsafe.SliceData(goSlice)          // *byteに変換し
		charPtr = (*C.char)(unsafe.Pointer(bytePtr)) // そこから (char *) に変換
		charLen = C.size_t(len(goSlice))             // サイズはスライスからそのまま取得
	)
	C.c_func2(charPtr, charLen)
}
```

- `unsafe.SliceData` は、Goのスライスの先頭要素へのポインタ (`*byte`) を返します。
- このポインタを `unsafe.Pointer` を経由してCの `*C.char` 型にキャストすることで、Cの関数に渡せるようになります。
- これにより、Goのメモリ領域をC側から直接読み取ることが可能になります。

### C言語とGo言語間の文字列の扱い

Cの文字列は通常NULL文字で終端されます。`unsafe.Slice`でGoスライスに変換した場合、このNULL文字もスライスの一部として含まれることがあります。
このサンプルでは、文字列を反転させる前に、NULL文字を考慮して実際のデータ長を計算しています。

```go
// NULL終端文字がある場合は減算して実データサイズとする
if dataLen > 0 && cSlice[dataLen-1] == 0 {
    dataLen--
}

// 実データ分をコピー
goSlice = make([]byte, dataLen)
copy(goSlice, cSlice[:dataLen])
```

逆に、GoからCへデータを渡す際は、C側がNULL終端文字列を期待していることを想定し、処理後のスライスにNULL文字を追加しています。

## 実行方法

プロジェクトのルートにある`Taskfile.yml`に実行コマンドが定義されています。以下のコマンドでサンプルを実行できます。

```sh
go run *.go
```

### 実行結果

```
[Go] helloworld
[C ] dlrowolleh
```
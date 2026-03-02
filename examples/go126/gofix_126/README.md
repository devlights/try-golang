# Go 1.26 の `go fix` 更新内容

## 変更の概要

`go fix` コマンドが、Go 1.10 の `go vet` と同様のパターンに倣い、**Go Analysis Framework**（`golang.org/x/tools/go/analysis`）を使用するよう刷新されました。

## 具体的な変更点

### 1. Analysis Frameworkへの移行

`go vet` でコードの診断（diagnostics）を提供するのと同じアナライザーが、`go fix` での修正提案・自動適用にも使用できるようになりました。

つまり、`go vet` と `go fix` が同一のアナライザー基盤を共有する設計になったということです。

### 2. 旧来のFixerの廃止

歴史的なfixerはすべて廃止されました。これらはいずれも**時代遅れ**（obsolete）のものでした。

### 3. 新しいアナライザー群への置き換え

旧来のfixerに代わり、**言語および標準ライブラリの新しい機能を活用するための修正を提案する**、新しいアナライザーのスイートが導入されました。

---

## 変更の意味・影響

| 観点 | 変更前 | 変更後 |
|---|---|---|
| 基盤 | 独自のfixer機構 | `go/analysis` Framework |
| `go vet` との関係 | 別々の仕組み | 同一アナライザーを共有可能 |
| 提供するfixer | 古い移行用（obsolete） | 新機能活用のための現代的なもの |
| カスタム拡張性 | 限定的 | Analysis Frameworkで統一 |


**参考情報：** https://go.dev/doc/go1.26

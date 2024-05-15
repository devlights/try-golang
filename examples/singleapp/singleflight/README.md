# これは何？

```golang.org/x/sync/singleflight``` のサンプルです。

singleflightパッケージは、重複した関数呼び出しを抑制するためのメカニズムを提供します。

このパッケージは、特に高価な操作や重複する操作が同時に複数のゴルーチンから要求される場合に有効です。

singleflightパッケージは、golang.org/x/sync/singleflightライブラリに含まれており、主に以下の機能を提供します。

- 重複呼び出しの抑制：同じキーに対する複数のリクエストが同時に発生した場合、最初のリクエストが完了するまで他のリクエストを待機させ、結果を共有します。
- 効率の向上：重複した操作を防ぐことで、サービスやデータベースへの不要な負荷を軽減します。
- シンプルなAPI：Group型を使用して、重複する操作を管理します。

Cache Stampedeなどが発生する可能性がある部分などで利用出来ます。

```Group.Forget()``` が存在するのがちょっとした違い。

## 参考情報

- https://pkg.go.dev/golang.org/x/sync/singleflight
- https://twitter.com/func25/status/1778770235316916427?t=39VDEN8c8WFp9fc-JGBABA&s=19
- https://zenn.dev/nkmrkz/articles/go-singleflight
- https://christina04.hatenablog.com/entry/go-singleflight
- https://pkg.go.dev/sync@go1.22.3#OnceValue
- https://cs.opensource.google/go/x/sync/+/refs/tags/v0.7.0:singleflight/singleflight.go
- https://blog.wu-boy.com/2024/02/how-to-reslove-the-hotspot-invalid-using-singleflight-en/
- https://www.codingexplorations.com/blog/understanding-singleflight-in-golang-a-solution-for-eliminating-redundant-work

/*
Package gotour -- Tour of Go (https://tour.golang.org/) の サンプルが配置されているパッケージです。

Go の プログラムは、パッケージで構成される。
規約により、パッケージ名はimportパスの最後の要素の名前となる。
プログラムは、必ず main パッケージから開始される。
Go では、一つのディレクトリ内に一つのパッケージしか含めることが出来ない。

パッケージコメントは、そのパッケージ配下のファイルのどこかに
記載されていればいいので、よくパターンとしては
 doc.go
というファイルを配置して、そこにパッケージコメントを記載する。

参考になる go.doc: https://golang.org/src/encoding/gob/doc.go

go doc の書き方については以下が分かりやすい

https://blog.golang.org/godoc-documenting-go-code
https://qiita.com/lufia/items/97acb391c26f967048f1
https://qiita.com/shibukawa/items/8c70fdd1972fad76a5ce
*/
package gotour

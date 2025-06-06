# これは何？

Goにてデバッグビルドするときと、リリースビルドするときによく利用するフラグについて。

## デバッグ

```sh
go build -gcflags "all=-N -l" -o debug main.go
```

### オプションの意味

```
      # gcflags の all=-N -l の意味 (goコンパイラに対しての指示) (go tool compile -help)
      #   all= は全てのパッケージが対象という意味
      #   -N   は最適化無効という意味 (No optimization)
      #   -l   はインライン化無効という意味 (No inlining)
```

## リリース

```sh
go build -ldflags "-s -w" -o release main.go
```

### オプションの意味

```
      # ldflags の -s -w の意味 (リンカに対しての指示) (go tool link -help)
      #    -s   はシンボルテーブル削除という意味
      #    -w   はDWARF情報削除という意味（デバッグ情報）
```

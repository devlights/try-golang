# これは何？

Goの書き方に慣れている人がC言語でプログラムを実装する際にやってしまいがちなことの一つ。

Goでは、「エスケープ分析」が存在するので関数ローカル変数のアドレスを戻り値として返却するコードを書いても

コンパイラがヒープに移動させてくれるため、問題にはならない。

が、C言語で同じことをすると当然うまく動かない。（コンパイル時に警告が出力されるので気づかないことは無いが）

C言語ではローカル変数のアドレスを返却した場合の挙動は「未定義」であるため、コンパイラによって挙動が変わる場合がある。

本サンプルをGitpod上で動作させたところ、gccの場合はコアダンプしたが、clangの場合はコアダンプはせず、値が上書きされて表示された。

Goの場合は呼び出しごとにヒープにエスケープされるので、狙った動作にはなる。

```sh
$ task
gcc (Ubuntu 11.4.0-1ubuntu1~22.04) 11.4.0
Ubuntu clang version 14.0.0-1ubuntu1.1
task: [compile-c] gcc -o ../app_gcc main.c
main.c: In function ‘getvalue’:
main.c:5:12: warning: function returns address of local variable [-Wreturn-local-addr]
    5 |     return &x;
      |            ^~
task: [compile-c] clang -o ../app_clang main.c
main.c:5:13: warning: address of stack memory associated with parameter 'x' returned [-Wreturn-stack-address]
    return &x;
            ^
1 warning generated.
task: [compile-go] go build -o app_go main.go
task: [run_gcc] sh -c './app_gcc | true'
Segmentation fault
task: [run_clang] ./app_clang
2,2
task: [run_go] ./app_go
1,2
```

## 参考情報

- https://www.dolthub.com/blog/2025-04-18-optimizing-heap-allocations/

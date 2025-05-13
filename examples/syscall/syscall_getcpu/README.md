# これは何？

syscall(2)を用いてgetcpu(2)を呼び出すサンプルです。

getcpu(2)システムコールはLinuxに存在しますが、glibc（GNU C ライブラリ）には対応するラッパー関数が提供されていません。
そのため、他の多くのシステムコールのように単純に関数として呼び出すことができません。
なので、syscall(2)を使って、glibcにラッパー関数が存在しないシステムコールを呼び出す必要があります。

```sh
$ task
task: [default] rm -f ./app
task: [default] go build -o app .
task: [default] ./app
CPU: 10, NUMA: 0
```

## C言語での実装

C言語では以下のようになります。

```c
#include <stdio.h>      // printf, perror
#include <stdlib.h>     // EXIT_SUCCESS, EXIT_FAILURE
#include <unistd.h>     // syscall
#include <sys/syscall.h> // SYS_getcpu
#include <errno.h>      // errno

/**
 * @brief getcpu(2)をsyscall(2)を使って呼び出すサンプルプログラム
 * 
 * このプログラムはsyscall(2)を使用してgetcpu(2)システムコールを直接呼び出し、
 * 現在のプロセスが実行されているCPU IDとNUMAノードを取得します。
 * 
 * getcpu(2)のプロトタイプ:
 * int getcpu(unsigned *cpu, unsigned *node, struct getcpu_cache *tcache);
 * 
 * @return 正常終了時は0、エラー時は-1
 */
int main(void)
{
    unsigned int cpu = 0;
    unsigned int node = 0;
    
    // getcpuシステムコールを呼び出す
    // 第3引数は通常NULLで良い（getcpu_cacheは非推奨）
    int result = syscall(SYS_getcpu, &cpu, &node, NULL);
    if (result == -1) {
        perror("syscall(SYS_getcpu) failed");
        return EXIT_FAILURE;
    }
    
    printf("現在のCPU ID: %u\n", cpu);
    printf("現在のNUMAノード: %u\n", node);
    
    return EXIT_SUCCESS;
}
```

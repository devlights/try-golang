# これは何？

[procfs](https://github.com/prometheus/procfs) を使って、/procファイルシステム上のCPU情報を取得するサンプルです。

```sh
$ task
task: [clean] rm -f ./app
task: [build] go build -o app .
task: [run] ./app
CPU-COUNT: 16
        [CPU-01] CoreId=0 ModelName="AMD EPYC 7B13", Processor=00
        [CPU-02] CoreId=1 ModelName="AMD EPYC 7B13", Processor=01
        [CPU-03] CoreId=2 ModelName="AMD EPYC 7B13", Processor=02
        [CPU-04] CoreId=3 ModelName="AMD EPYC 7B13", Processor=03
        [CPU-05] CoreId=4 ModelName="AMD EPYC 7B13", Processor=04
        [CPU-06] CoreId=5 ModelName="AMD EPYC 7B13", Processor=05
        [CPU-07] CoreId=6 ModelName="AMD EPYC 7B13", Processor=06
        [CPU-08] CoreId=7 ModelName="AMD EPYC 7B13", Processor=07
        [CPU-09] CoreId=0 ModelName="AMD EPYC 7B13", Processor=08
        [CPU-10] CoreId=1 ModelName="AMD EPYC 7B13", Processor=09
        [CPU-11] CoreId=2 ModelName="AMD EPYC 7B13", Processor=10
        [CPU-12] CoreId=3 ModelName="AMD EPYC 7B13", Processor=11
        [CPU-13] CoreId=4 ModelName="AMD EPYC 7B13", Processor=12
        [CPU-14] CoreId=5 ModelName="AMD EPYC 7B13", Processor=13
        [CPU-15] CoreId=6 ModelName="AMD EPYC 7B13", Processor=14
        [CPU-16] CoreId=7 ModelName="AMD EPYC 7B13", Processor=15
```
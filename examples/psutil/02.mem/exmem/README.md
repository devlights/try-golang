# これは何？

[gopsutil](https://github.com/shirou/gopsutil)を用いてメモリ消費量を取得するサンプルです。

v4.24.5から追加された Ex構造体 を利用するサンプルです。

```sh
$ task
task: [clean] rm -f ./app
task: [build] go build -o app .
task: [run] ./app
07:40:14 [NORMAL] {"total":67421265920,"available":52788764672,"used":13726756864,"usedPercent":20.359684257913145,"free":31725797376,"active":6271885312,"inactive":21599387648,"wired":0,"laundry":0,"buffers":1032192,"cached":21967679488,"writeBack":0,"dirty":3129344,"writeBackTmp":0,"shared":158375936,"slab":5851881472,"sreclaimable":3740106752,"sunreclaim":2111774720,"pageTables":174186496,"swapCached":22228992,"commitLimit":436363812864,"committedAS":31354429440,"highTotal":0,"highFree":0,"lowTotal":0,"lowFree":0,"swapTotal":402653179904,"swapFree":402577944576,"mapped":1446465536,"vmallocTotal":35184372087808,"vmallocUsed":204656640,"vmallocChunk":0,"hugePagesTotal":0,"hugePagesFree":0,"hugePagesRsvd":0,"hugePagesSurp":0,"hugePageSize":2097152,"anonHugePages":12582912}
07:40:14 [EX    ] {"activefile":5806538752,"inactivefile":12259991552,"activeanon":465346560,"inactiveanon":9339396096,"unevictable":87674880}
```

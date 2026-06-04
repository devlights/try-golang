# 実行結果

## Gitpod

```sh
$ lscpu | grep -E '^(CPU\(s\):|Model name:).*$'
CPU(s):                               16
Model name:                           AMD EPYC 7B13

$ task
task: [default] rm -f ./app
task: [default] goimports -w main.go
task: [default] go build -o app main.go
task: [default] time ./app -loop 500000 -inch 0 -outch 0
numWorkers=16
done

real    0m0.537s
user    0m0.000s
sys     0m0.000s
task: [default] time ./app -loop 500000 -inch 500000 -outch 0
numWorkers=16
done

real    0m0.419s
user    0m0.000s
sys     0m0.000s
task: [default] time ./app -loop 500000 -inch 500000 -outch 500000
numWorkers=16
done

real    0m0.284s
user    0m0.000s
sys     0m0.000s
```

## Github Codespaces (CPU=2)

```sh
$ lscpu | grep -E '^(CPU\(s\):|Model name:).*$'
CPU(s):                               2
Model name:                           AMD EPYC 7763 64-Core Processor

$ task
task: [default] rm -f ./app
task: [default] goimports -w main.go
task: [default] go build -o app main.go
task: [default] time ./app -loop 500000 -inch 0 -outch 0
numWorkers=2
done

real    0m0.605s
user    0m0.000s
sys     0m0.000s
task: [default] time ./app -loop 500000 -inch 500000 -outch 0
numWorkers=2
done

real    0m0.546s
user    0m0.000s
sys     0m0.000s
task: [default] time ./app -loop 500000 -inch 500000 -outch 500000
numWorkers=2
done

real    0m0.269s
user    0m0.000s
sys     0m0.000s
```

## Chromebook

```sh
$ lscpu | grep -E '^(CPU\(s\):|Model name:).*$'
CPU(s):                               8
Model name:                           Kryo-4XX-Silver
Model name:                           Kryo-4XX-Gold

$ task
task: [default] rm -f ./app
task: [default] goimports -w main.go
task: [default] go build -o app main.go
task: [default] time ./app -loop 500000 -inch 0 -outch 0
numWorkers=8
done

real    0m1.264s
user    0m0.000s
sys     0m0.000s
task: [default] time ./app -loop 500000 -inch 500000 -outch 0
numWorkers=8
done

real    0m1.127s
user    0m0.000s
sys     0m0.000s
task: [default] time ./app -loop 500000 -inch 500000 -outch 500000
numWorkers=8
done

real    0m0.403s
user    0m0.000s
sys     0m0.000s
```

---

# GitpodとGithub CodespacesのCPUについて (by Claude 3.7 Sonnet)

## AMD EPYC 7B13と7763 64-Coreプロセッサの性能比較

### 基本仕様の比較

#### AMD EPYC 7B13
- **コア/スレッド**: 64コア/128スレッド AMD EPYC 7B13は64コアのプロセッサで、マルチスレッド対応により128スレッドを実現しています。
- **基本クロック**: 2.2GHz (2.3GHz情報もあり) 7B13のCPUベース周波数は2200MHzとなっています。 OpenBenchmarkingの情報によると2.3GHzのクロック速度とされています。
- **最大ブースト**: 3.5GHz 最大周波数は3500MHzです。
- **L3キャッシュ**: 256MB L3キャッシュは256MBを搭載しています。
- **ソケットタイプ**: SP3 SP3ソケットを使用しています。
- **マイクロアーキテクチャ**: Zen 3 AMD EPYC 7B13は「Zen 3」プロセッサです。
- **TDP**: 明確な記載はありませんが、同世代の64コアEPYCプロセッサから推定すると、225-280W程度と考えられます
- **PassMark CPU Mark**: 77,460 ベンチマークテストでは77,460のスコアを記録しています。
- **シングルスレッドスコア**: 2,564 シングルスレッドレーティングは2,564です。

#### AMD EPYC 7763
- **コア/スレッド**: 64コア/128スレッド EPYC 7763は64コア128スレッドのプロセッサです。
- **基本クロック**: 2.45GHz 基本周波数は2.45GHzです。
- **最大ブースト**: 3.5GHz ターボブースト時の最大周波数は3.5GHzに達します。
- **L3キャッシュ**: 256MB EPYC 7763は256MBのL3キャッシュを搭載しています。
- **ソケットタイプ**: SP3 SP3ソケットに対応しています。
- **マイクロアーキテクチャ**: Zen 3 Zen 3マイクロアーキテクチャに基づいています。
- **TDP**: 280W TDPは280Wで、かなりの電力を消費します。
- **PassMark CPU Mark**: 84,522〜84,591 ベンチマークスコアはPassMark CPU Markテストで84,522〜84,591を記録しています。
- **シングルスレッドスコア**: 2,518〜2,525 シングルスレッドレーティングは2,518〜2,525です。

### パフォーマンス比較

AMD EPYC 7B13はEPYC 7763よりも新しいモデルですが、マルチスレッド性能では7763が約8%高速です。一方、シングルスレッドの性能はほぼ同等となっています。

ベンチマークスコアを具体的に比較すると：
- **マルチスレッド性能**: EPYC 7763が84,522〜84,591に対し、7B13は77,460であり、約8-9%の差があります PassMarkのCPU Markスコアではこの差が確認できます。
- **シングルスレッド性能**: EPYC 7763が2,518〜2,525に対し、7B13は2,564と、わずかに7B13が上回っています シングルスレッドでは僅かながら7B13の方が優れています。

### 用途と市場ポジション

EPYC 7763：
7763はAMD EPYCシリーズの高性能モデルとして位置づけられており、最高レベルのパフォーマンスと機能を提供するために設計されています。より低いコア数と消費電力の「メインストリーム」モデルと比較して、トップビンのパフォーマンスを実現するためのモデルです。

このような高価なCPUの市場は、システムコストの大部分がアクセラレータ（GPUなど）に占められる環境です。本質的に、可能な限り最高のCPUを搭載することで、アクセラレータが最大限の速度で動作することを保証します。

EPYC 7B13：
7B13はGoogle Cloud Platform向けの特別モデルとされており eBayの出品情報によると「7B13 Google Cloud Platform 100-000000335」と記載されています。、クラウドプロバイダー向けに最適化されている可能性があります。

### 電力効率

EPYC 7763はTDP 280Wと非常に電力を消費するため、高品質な冷却システムが必要です。

7B13のTDPについては具体的な情報が見つかりませんでしたが、パフォーマンスが若干低いことを考慮すると、7763よりも若干低い可能性があります。

### まとめ

AMD EPYC 7B13と7763を比較すると、以下のような特徴が見えてきます：

1. **基本アーキテクチャ**: 両プロセッサともZen 3アーキテクチャを採用し、64コア/128スレッド、256MB L3キャッシュという強力な仕様を持っています。

2. **クロック周波数**: EPYC 7763の方が基本クロックが僅かに高く（2.45GHz vs 2.2-2.3GHz）、最大ブースト周波数は両方とも3.5GHzです。

3. **性能差**: 
   - マルチスレッド性能では7763が約8-9%優れています
   - シングルスレッド性能では7B13がわずかに上回っています

4. **用途**: 
   - EPYC 7763はハイエンドサーバー向けの高性能モデルとして位置づけられています
   - EPYC 7B13はGoogle Cloud Platform向けの特別モデルとされ、クラウド環境向けに最適化されている可能性があります

5. **消費電力**:
   - EPYC 7763は280WのTDPを持ち、非常に電力を消費します
   - 7B13のTDPは明確ではありませんが、パフォーマンスから推測するとやや低い可能性があります

結論として、同じZen 3アーキテクチャと64コアを持つ両プロセッサは非常に高性能ですが、若干の差異があります。マルチスレッド処理が最重要の環境では7763が有利ですが、クラウド環境や電力効率を重視する場合は7B13も魅力的な選択肢となるでしょう。

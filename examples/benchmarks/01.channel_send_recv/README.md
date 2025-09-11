# これは何？

チャネルの純粋な送受信がどれくらいの速度で動作するのかについてのベンチマークです。

## 実行例

```sh
$ task run
task: [run] go test -bench . -count 10 | tee bench.out
goos: linux
goarch: amd64
pkg: github.com/devlights/try-golang/examples/benchmarks/channel_send_recv
cpu: AMD EPYC 7B13
BenchmarkChanSync/100-16 	  279804	      4031 ns/op	    1032 B/op	       2 allocs/op
BenchmarkChanSync/100-16 	  302364	      4064 ns/op	    1032 B/op	       2 allocs/op
BenchmarkChanSync/100-16 	  274219	      4072 ns/op	    1032 B/op	       2 allocs/op
BenchmarkChanSync/100-16 	  290083	      4071 ns/op	    1032 B/op	       2 allocs/op
BenchmarkChanSync/100-16 	  294525	      4030 ns/op	    1032 B/op	       2 allocs/op
BenchmarkChanSync/100-16 	  287929	      4288 ns/op	    1032 B/op	       2 allocs/op
BenchmarkChanSync/100-16 	  276304	      4064 ns/op	    1032 B/op	       2 allocs/op
BenchmarkChanSync/100-16 	  292827	      4047 ns/op	    1032 B/op	       2 allocs/op
BenchmarkChanSync/100-16 	  295461	      3967 ns/op	    1032 B/op	       2 allocs/op
BenchmarkChanSync/100-16 	  291420	      4059 ns/op	    1032 B/op	       2 allocs/op
BenchmarkChanSync/1000-16         	   30254	     39316 ns/op	    8205 B/op	       2 allocs/op
BenchmarkChanSync/1000-16         	   29754	     39928 ns/op	    8205 B/op	       2 allocs/op
BenchmarkChanSync/1000-16         	   30517	     39586 ns/op	    8205 B/op	       2 allocs/op
BenchmarkChanSync/1000-16         	   30960	     39365 ns/op	    8205 B/op	       2 allocs/op
BenchmarkChanSync/1000-16         	   27843	     41077 ns/op	    8205 B/op	       2 allocs/op
BenchmarkChanSync/1000-16         	   30912	     38975 ns/op	    8205 B/op	       2 allocs/op
BenchmarkChanSync/1000-16         	   30784	     38748 ns/op	    8205 B/op	       2 allocs/op
BenchmarkChanSync/1000-16         	   30766	     38676 ns/op	    8205 B/op	       2 allocs/op
BenchmarkChanSync/1000-16         	   29863	     39719 ns/op	    8205 B/op	       2 allocs/op
BenchmarkChanSync/1000-16         	   31249	     38666 ns/op	    8205 B/op	       2 allocs/op
BenchmarkChanSync/10000-16        	    2983	    389806 ns/op	   81983 B/op	       2 allocs/op
BenchmarkChanSync/10000-16        	    3055	    384617 ns/op	   81980 B/op	       2 allocs/op
BenchmarkChanSync/10000-16        	    3229	    391278 ns/op	   81981 B/op	       2 allocs/op
BenchmarkChanSync/10000-16        	    3098	    387064 ns/op	   81981 B/op	       2 allocs/op
BenchmarkChanSync/10000-16        	    3178	    383327 ns/op	   81981 B/op	       2 allocs/op
BenchmarkChanSync/10000-16        	    3002	    386414 ns/op	   81981 B/op	       2 allocs/op
BenchmarkChanSync/10000-16        	    3109	    383222 ns/op	   81982 B/op	       2 allocs/op
BenchmarkChanSync/10000-16        	    2918	    389171 ns/op	   81981 B/op	       2 allocs/op
BenchmarkChanSync/10000-16        	    3013	    387681 ns/op	   81981 B/op	       2 allocs/op
BenchmarkChanSync/10000-16        	    3048	    383782 ns/op	   81981 B/op	       2 allocs/op
BenchmarkChanSync/100000-16       	     316	   3972062 ns/op	  803420 B/op	       2 allocs/op
BenchmarkChanSync/100000-16       	     286	   3940878 ns/op	  803423 B/op	       2 allocs/op
BenchmarkChanSync/100000-16       	     312	   3853476 ns/op	  803411 B/op	       2 allocs/op
BenchmarkChanSync/100000-16       	     312	   3890827 ns/op	  803413 B/op	       2 allocs/op
BenchmarkChanSync/100000-16       	     315	   3846943 ns/op	  803424 B/op	       2 allocs/op
BenchmarkChanSync/100000-16       	     306	   3919488 ns/op	  803423 B/op	       2 allocs/op
BenchmarkChanSync/100000-16       	     307	   3916240 ns/op	  803419 B/op	       2 allocs/op
BenchmarkChanSync/100000-16       	     294	   3941979 ns/op	  803423 B/op	       2 allocs/op
BenchmarkChanSync/100000-16       	     309	   3969874 ns/op	  803425 B/op	       2 allocs/op
BenchmarkChanSync/100000-16       	     304	   3870358 ns/op	  803413 B/op	       2 allocs/op
BenchmarkChanSync/1000000-16      	      28	  38731244 ns/op	 8004792 B/op	       3 allocs/op
BenchmarkChanSync/1000000-16      	      31	  38448564 ns/op	 8004835 B/op	       3 allocs/op
BenchmarkChanSync/1000000-16      	      28	  38591415 ns/op	 8004808 B/op	       3 allocs/op
BenchmarkChanSync/1000000-16      	      31	  38557124 ns/op	 8004810 B/op	       3 allocs/op
BenchmarkChanSync/1000000-16      	      27	  39037175 ns/op	 8004824 B/op	       3 allocs/op
BenchmarkChanSync/1000000-16      	      30	  39187820 ns/op	 8004822 B/op	       3 allocs/op
BenchmarkChanSync/1000000-16      	      30	  38997505 ns/op	 8004781 B/op	       3 allocs/op
BenchmarkChanSync/1000000-16      	      31	  38464896 ns/op	 8004857 B/op	       3 allocs/op
BenchmarkChanSync/1000000-16      	      30	  39552623 ns/op	 8004766 B/op	       3 allocs/op
BenchmarkChanSync/1000000-16      	      26	  38956326 ns/op	 8004833 B/op	       3 allocs/op
BenchmarkChanAsync/100-16         	  186831	      5874 ns/op	    1161 B/op	       7 allocs/op
BenchmarkChanAsync/100-16         	  201355	      5977 ns/op	    1161 B/op	       7 allocs/op
BenchmarkChanAsync/100-16         	  203440	      6075 ns/op	    1160 B/op	       7 allocs/op
BenchmarkChanAsync/100-16         	  205980	      6014 ns/op	    1160 B/op	       7 allocs/op
BenchmarkChanAsync/100-16         	  204482	      6194 ns/op	    1160 B/op	       7 allocs/op
BenchmarkChanAsync/100-16         	  190962	      5961 ns/op	    1160 B/op	       7 allocs/op
BenchmarkChanAsync/100-16         	  197955	      6332 ns/op	    1160 B/op	       7 allocs/op
BenchmarkChanAsync/100-16         	  180174	      6139 ns/op	    1160 B/op	       7 allocs/op
BenchmarkChanAsync/100-16         	  202866	      6156 ns/op	    1160 B/op	       7 allocs/op
BenchmarkChanAsync/100-16         	  202098	      6012 ns/op	    1160 B/op	       7 allocs/op
BenchmarkChanAsync/1000-16        	   26412	     45228 ns/op	    8334 B/op	       7 allocs/op
BenchmarkChanAsync/1000-16        	   24825	     48751 ns/op	    8334 B/op	       7 allocs/op
BenchmarkChanAsync/1000-16        	   27230	     44634 ns/op	    8334 B/op	       7 allocs/op
BenchmarkChanAsync/1000-16        	   26161	     45000 ns/op	    8334 B/op	       7 allocs/op
BenchmarkChanAsync/1000-16        	   27258	     45609 ns/op	    8334 B/op	       7 allocs/op
BenchmarkChanAsync/1000-16        	   24712	     46048 ns/op	    8334 B/op	       7 allocs/op
BenchmarkChanAsync/1000-16        	   26526	     45107 ns/op	    8334 B/op	       7 allocs/op
BenchmarkChanAsync/1000-16        	   26499	     45045 ns/op	    8335 B/op	       7 allocs/op
BenchmarkChanAsync/1000-16        	   24843	     47015 ns/op	    8335 B/op	       7 allocs/op
BenchmarkChanAsync/1000-16        	   26824	     44760 ns/op	    8334 B/op	       7 allocs/op
BenchmarkChanAsync/10000-16       	    2098	    536619 ns/op	   82120 B/op	       7 allocs/op
BenchmarkChanAsync/10000-16       	    2265	    532042 ns/op	   82120 B/op	       7 allocs/op
BenchmarkChanAsync/10000-16       	    2313	    534333 ns/op	   82119 B/op	       7 allocs/op
BenchmarkChanAsync/10000-16       	    2124	    540292 ns/op	   82118 B/op	       7 allocs/op
BenchmarkChanAsync/10000-16       	    2268	    527942 ns/op	   82119 B/op	       7 allocs/op
BenchmarkChanAsync/10000-16       	    2200	    539408 ns/op	   82119 B/op	       7 allocs/op
BenchmarkChanAsync/10000-16       	    2232	    536863 ns/op	   82120 B/op	       7 allocs/op
BenchmarkChanAsync/10000-16       	    2238	    539517 ns/op	   82119 B/op	       7 allocs/op
BenchmarkChanAsync/10000-16       	    2214	    528823 ns/op	   82119 B/op	       7 allocs/op
BenchmarkChanAsync/10000-16       	    2320	    528874 ns/op	   82120 B/op	       7 allocs/op
BenchmarkChanAsync/100000-16      	     200	   5951932 ns/op	  803600 B/op	       8 allocs/op
BenchmarkChanAsync/100000-16      	     194	   6225372 ns/op	  803622 B/op	       8 allocs/op
BenchmarkChanAsync/100000-16      	     200	   5940656 ns/op	  803594 B/op	       8 allocs/op
BenchmarkChanAsync/100000-16      	     199	   6100676 ns/op	  803604 B/op	       8 allocs/op
BenchmarkChanAsync/100000-16      	     198	   6042795 ns/op	  803597 B/op	       8 allocs/op
BenchmarkChanAsync/100000-16      	     196	   6151063 ns/op	  803593 B/op	       8 allocs/op
BenchmarkChanAsync/100000-16      	     194	   6200071 ns/op	  803609 B/op	       8 allocs/op
BenchmarkChanAsync/100000-16      	     196	   6236949 ns/op	  803610 B/op	       8 allocs/op
BenchmarkChanAsync/100000-16      	     195	   6140092 ns/op	  803647 B/op	       8 allocs/op
BenchmarkChanAsync/100000-16      	     190	   6236760 ns/op	  803595 B/op	       8 allocs/op
BenchmarkChanAsync/1000000-16     	      19	  60881378 ns/op	 8005094 B/op	       9 allocs/op
BenchmarkChanAsync/1000000-16     	      19	  60285705 ns/op	 8005045 B/op	       9 allocs/op
BenchmarkChanAsync/1000000-16     	      19	  58995496 ns/op	 8005061 B/op	       9 allocs/op
BenchmarkChanAsync/1000000-16     	      19	  61019145 ns/op	 8005164 B/op	      10 allocs/op
BenchmarkChanAsync/1000000-16     	      18	  63189106 ns/op	 8005026 B/op	       9 allocs/op
BenchmarkChanAsync/1000000-16     	      20	  60307970 ns/op	 8005048 B/op	       9 allocs/op
BenchmarkChanAsync/1000000-16     	      18	  61124251 ns/op	 8005091 B/op	       9 allocs/op
BenchmarkChanAsync/1000000-16     	      18	  61482151 ns/op	 8005060 B/op	       9 allocs/op
BenchmarkChanAsync/1000000-16     	      20	  60296411 ns/op	 8005076 B/op	       9 allocs/op
BenchmarkChanAsync/1000000-16     	      20	  60038558 ns/op	 8005068 B/op	       9 allocs/op
PASS
ok  	github.com/devlights/try-golang/examples/benchmarks/channel_send_recv	118.281s
```

## benchstatでの表示

```sh
$ task stat
task: [stat] benchstat bench.out
goos: linux
goarch: amd64
pkg: github.com/devlights/try-golang/examples/benchmarks/channel_send_recv
cpu: AMD EPYC 7B13
                     │  bench.out  │
                     │   sec/op    │
ChanSync/100-16        4.062µ ± 1%
ChanSync/1000-16       39.34µ ± 2%
ChanSync/10000-16      386.7µ ± 1%
ChanSync/100000-16     3.918m ± 2%
ChanSync/1000000-16    38.84m ± 1%
ChanAsync/100-16       6.045µ ± 2%
ChanAsync/1000-16      45.17µ ± 4%
ChanAsync/10000-16     535.5µ ± 1%
ChanAsync/100000-16    6.146m ± 3%
ChanAsync/1000000-16   60.59m ± 1%
geomean                468.7µ

                     │  bench.out   │
                     │     B/op     │
ChanSync/100-16        1.008Ki ± 0%
ChanSync/1000-16       8.013Ki ± 0%
ChanSync/10000-16      80.06Ki ± 0%
ChanSync/100000-16     784.6Ki ± 0%
ChanSync/1000000-16    7.634Mi ± 0%
ChanAsync/100-16       1.133Ki ± 0%
ChanAsync/1000-16      8.139Ki ± 0%
ChanAsync/10000-16     80.19Ki ± 0%
ChanAsync/100000-16    784.8Ki ± 0%
ChanAsync/1000000-16   7.634Mi ± 0%
geomean                84.24Ki

                     │ bench.out  │
                     │ allocs/op  │
ChanSync/100-16        2.000 ± 0%
ChanSync/1000-16       2.000 ± 0%
ChanSync/10000-16      2.000 ± 0%
ChanSync/100000-16     2.000 ± 0%
ChanSync/1000000-16    3.000 ± 0%
ChanAsync/100-16       7.000 ± 0%
ChanAsync/1000-16      7.000 ± 0%
ChanAsync/10000-16     7.000 ± 0%
ChanAsync/100000-16    8.000 ± 0%
ChanAsync/1000000-16   9.000 ± 0%
geomean                4.049
```
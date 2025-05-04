# Result

## Gitpod

```sh
$ lscpu | grep 'CPU(s):'
CPU(s):                               16
NUMA node0 CPU(s):                    0-15

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


## Chromebook

```sh
$ lscpu | grep 'CPU(s):'
CPU(s):                               8

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

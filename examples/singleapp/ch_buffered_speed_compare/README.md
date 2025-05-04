# Result

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

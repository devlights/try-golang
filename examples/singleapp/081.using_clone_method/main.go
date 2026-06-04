package main

import (
	"flag"
	"io"
	"log"
	"os/exec"
	"runtime"
	"strings"
	"unsafe"
)

const (
	NumItems = 1000
	SHELL    = "/bin/bash"
)

var (
	store = make([]string, NumItems)
)

func init() {
	log.SetFlags(0)
}

func mem(prefix string) {
	var (
		m = runtime.MemStats{}
	)

	runtime.ReadMemStats(&m)
	log.Printf("[%s]\t%8d\t%8d\n", prefix, m.HeapAlloc, m.HeapObjects)
}

func gen() []string {
	var (
		l = make([]string, NumItems)
	)

	for i := 0; i < NumItems; i++ {
		output, _ := exec.Command(SHELL, "-c", "openssl rand -base64 4096 | tr -d '\n'").Output()
		l[i] = unsafe.String(&output[0], len(output))
	}

	return l
}

func main() {
	log.Println("Title       \tHeapAlloc\tHeapObjects")
	mem("start     ")

	var (
		use = flag.Bool("use", false, "Use strings.Clone()")
	)
	flag.Parse()

	var (
		l = gen()
	)
	mem("gen       ")

	for i := 0; i < NumItems; i++ {
		storeValue := l[i][:5]

		if *use {
			store[i] = strings.Clone(storeValue)
		} else {
			store[i] = storeValue
		}
	}
	mem("store     ")

	runtime.GC()

	for i, v := range store {
		if i%200 == 0 {
			runtime.GC()
			mem("checkpoint")
		}

		io.Discard.Write(unsafe.Slice(unsafe.StringData(v), len(v)))
	}

	/*
		$ task
		task: [build] go build -o app main.go
		task: [run-not-use-clone] ./app
		Title           HeapAlloc       HeapObjects
		[start     ]      192792             144
		[gen       ]    11482528            4576
		[store     ]    11487008            4588
		[checkpoint]     8471296            1363
		[checkpoint]     8475728            1372
		[checkpoint]     8475728            1372
		[checkpoint]     8475728            1372
		[checkpoint]     8475736            1373
		task: [run-use-clone] ./app -use
		Title           HeapAlloc       HeapObjects
		[start     ]      192824             144
		[gen       ]    11497632            4607
		[store     ]    11507440            4952
		[checkpoint]      296112             724
		[checkpoint]      300536             732
		[checkpoint]      300536             732
		[checkpoint]      300544             733
		[checkpoint]      300544             733
	*/
}

package main

import (
	"io"
	"os"
	"testing"
)

var (
	fname string
	buf   = make([]byte, 10)
)

func setup() {
	f, err := os.CreateTemp("", "trygolang-tmp-")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = f.Write([]byte("helloworld\n"))
	if err != nil {
		panic(err)
	}

	fname = f.Name()
}

func teardown() {
	err := os.Remove(fname)
	if err != nil {
		panic(err)
	}
}

func BenchmarkOsOpenEvery(b *testing.B) {
	setup()
	b.Cleanup(teardown)
	b.ResetTimer()

	for range b.N {
		f, err := os.Open(fname)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		clear(buf)

		_, err = f.Read(buf)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkOsOpenKeep(b *testing.B) {
	setup()
	b.Cleanup(teardown)
	b.ResetTimer()

	f, err := os.Open(fname)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	for range b.N {
		_, err = f.Seek(0, io.SeekStart)
		if err != nil {
			panic(err)
		}

		clear(buf)

		_, err = f.Read(buf)
		if err != nil {
			panic(err)
		}
	}
}

package main

import (
	"io/fs"
	"os"
	"path/filepath"
	"testing"
)

const (
	dir = "../../../"
)

func TestMain(m *testing.M) {
	err := os.Chdir(dir)
	if err != nil {
		os.Exit(1)
	}

	ret := m.Run()

	os.Exit(ret)
}

func BenchmarkWalk(b *testing.B) {
	for i := 0; i < b.N; i++ {
		filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
			return nil
		})
	}
}

func BenchmarkWalkDir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		filepath.WalkDir(".", func(path string, info fs.DirEntry, err error) error {
			return nil
		})
	}
}

package shm

import (
	"os"
	"path/filepath"
	"sync"
)

type (
	shm struct {
		f  *os.File
		mu sync.RWMutex
	}
)

func New(name string) (*shm, error) {
	f, err := os.OpenFile(filepath.Join("/dev/shm", name), os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		return nil, err
	}

	return &shm{f, sync.RWMutex{}}, nil
}

func Open(name string) (*shm, error) {
	f, err := os.Open(filepath.Join("/dev/shm", name))
	if err != nil {
		return nil, err
	}

	return &shm{f, sync.RWMutex{}}, nil
}

func (me *shm) Close() error {
	if me == nil {
		return nil
	}

	if me.f == nil {
		return nil
	}

	return me.f.Close()
}

func (me *shm) Unlink() error {
	if me == nil {
		return nil
	}

	if me.f == nil {
		return nil
	}

	me.mu.Lock()
	defer me.mu.Unlock()

	stat, err := me.f.Stat()
	if err != nil {
		return err
	}

	if err := me.Close(); err != nil {
		return err
	}

	if err := os.Remove(filepath.Join("/dev/shm", stat.Name())); err != nil {
		return err
	}
	me.f = nil

	return nil
}

func (me *shm) Read(b []byte) (int, error) {
	if me == nil {
		return 0, nil
	}

	if me.f == nil {
		return 0, nil
	}

	me.mu.RLock()
	defer me.mu.RUnlock()

	if err := me.f.Sync(); err != nil {
		return 0, err
	}

	return me.f.Read(b)
}

func (me *shm) Write(b []byte) (int, error) {
	if me == nil {
		return 0, nil
	}

	if me.f == nil {
		return 0, nil
	}

	me.mu.Lock()
	defer me.mu.Unlock()

	n, err := me.f.Write(b)
	if err != nil {
		return n, err
	}

	if err := me.f.Sync(); err != nil {
		return n, err
	}

	return n, nil
}

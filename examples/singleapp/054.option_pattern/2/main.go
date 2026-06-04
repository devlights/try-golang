package main

import (
	"errors"
	"fmt"
	"time"
)

type (
	Server struct {
		Addr        string
		SendTimeout time.Duration
	}

	options struct {
		sendTimeout time.Duration
	}

	Option func(o *options) error
)

func NewServer(addr string, opts ...Option) (*Server, error) {
	var options options
	for _, opt := range opts {
		err := opt(&options)
		if err != nil {
			return nil, err
		}
	}

	s := new(Server)
	s.Addr = addr
	s.SendTimeout = options.sendTimeout

	return s, nil
}

func WithSendTimeout(v time.Duration) Option {
	return func(o *options) error {
		if v < 1*time.Second {
			return errors.New("value should be greater than 1")
		}
		o.sendTimeout = v
		return nil
	}
}

func main() {
	s1, err := NewServer(":8888")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", s1)

	s2, err := NewServer(":8888", WithSendTimeout(0*time.Second))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", s2)

	s3, err := NewServer(":8888", WithSendTimeout(3*time.Second))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", s3)
}

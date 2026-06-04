package config

import (
	"fmt"
	"time"
)

type (
	Config struct {
		Addr        string
		Port        int
		RecvTimeout time.Duration
		SendTimeout time.Duration
	}

	// Option Pattern を実現するために用意
	Option func(c *Config)
)

func New(addr string, port int, options ...Option) *Config {
	c := new(Config)

	c.Addr = addr
	c.Port = port

	// Option Pattern
	for _, opt := range options {
		opt(c)
	}

	return c
}

func (c *Config) String() string {
	return fmt.Sprintf(
		"addr=%v:%v\trecvTimeout=%v\tsendTimeout=%v",
		c.Addr,
		c.Port,
		c.RecvTimeout,
		c.SendTimeout)
}

// WithRecvTimeout is implemented Option Pattern
func WithRecvTimeout(v time.Duration) Option {
	return func(c *Config) {
		c.RecvTimeout = v
	}
}

// WithSendTimeout is implemented Option Pattern
func WithSendTimeout(v time.Duration) Option {
	return func(c *Config) {
		c.SendTimeout = v
	}
}

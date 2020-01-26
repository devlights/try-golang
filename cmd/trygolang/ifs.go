package main

type (
	Command interface {
		Run() error
	}
)

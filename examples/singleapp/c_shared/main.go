package main

import "C"
import (
	"log"
)

//export GoAdd
func GoAdd(x, y int) int {
	return x + y
}

func init() {
	log.SetFlags(0)
	log.Println("[FROM GOLANG] library loaded!")
}

func main() {
}

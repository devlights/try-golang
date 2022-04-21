package main

import (
	"fmt"
	"log"
	"os/user"
)

func main() {
	u, err := user.Current()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Username: %s\n", u.Username)
}

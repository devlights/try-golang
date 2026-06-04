package main

import (
	"context"
	"fmt"
	"log"
)

func init() {
	log.SetFlags(0)
}

func main() {
	var (
		rootCtx          = context.Background()
		mainCtx, mainCxl = context.WithCancel(rootCtx)
		err              error
	)
	defer mainCxl()

	if err = run(mainCtx); err != nil {
		log.Fatal(err)
	}
}

func run(_ context.Context) error {
	var (
		circular = NewCircular[string](5)
	)

	//
	// Add
	//
	title("Add")
	{
		for _, r := range "helloworld" {
			circular.Add(string(r))
			log.Println(circular)
		}
	}

	//
	// Iterate (slice)
	//
	title("Iterate (slice)")
	{
		for i, v := range circular.ToSlice() {
			log.Printf("%02d: %v", i, v)
		}
	}

	//
	// Iterate (node)
	//
	title("Iterate (node)")
	{
		if circular.Head != nil {
			for n := circular.Head; ; n = n.Next {
				log.Println(n)

				if n == circular.Tail {
					break
				}
			}
		}
	}

	//
	// Delete
	//
	title("Delete")
	{
		var (
			fn = func(v1, v2 string) bool {
				return v1 == v2
			}
			deletes = []string{"d", "w", "r"}
		)
		for _, d := range deletes {
			if ok := circular.Delete(d, fn); !ok {
				return fmt.Errorf("Delete() returns false (%v)", d)
			}

			log.Println(circular)
		}
	}

	return nil
}

func title(m string) {
	log.Printf("----------[%s]----------", m)
}

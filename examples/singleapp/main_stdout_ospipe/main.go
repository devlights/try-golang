package main

import (
	"flag"
	"fmt"
	"strings"
)

type (
	vars []string
)

func (me *vars) String() string {
	return fmt.Sprint(*me)
}

func (me *vars) Set(v string) error {
	*me = append(*me, v)
	return nil
}

var (
	_ flag.Value = (*vars)(nil)
)

func main() {
	var (
		vs vars
	)
	flag.Var(&vs, "v", "values")
	flag.Parse()

	fmt.Println(strings.Join(vs, ","))
}

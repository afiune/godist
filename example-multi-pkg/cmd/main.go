package main

import (
	"fmt"

	"github.com/afiune/godist/example-multi-pkg/dist"
)

func main() {
	fmt.Printf("Global Foo variable: '%s'\n", dist.Foo)
	fmt.Printf("Local Exec variable: '%s'\n", dist.Exec)
}

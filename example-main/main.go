package main

import "fmt"

func main() {
	fmt.Printf("%s:\n", Exec)
	fmt.Printf(" * Global product name: '%s'\n", ServerProduct)
	fmt.Printf(" * Local variable: '%s'\n", LocalVar)
}

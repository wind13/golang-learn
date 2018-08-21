package main

import (
 "flag"
 "fmt"
)

func main() {
	// Here name is a cursor(*string) not a string.
	var name = flag.String("name", "everyone", "The greeting object.")
 	fmt.Printf("1. %v!, %v!\n", name, *name)
	flag.Parse()
 	fmt.Printf("2. %v!, %v!\n", name, *name)
 	fmt.Printf("Hello, %v!\n", *name)
}
package main

import (
 "flag"
 "fmt"
)

func main() {
 	var name string // [1]
 	flag.StringVar(&name, "name", "everyone", "The greeting object.") // [2]
 	fmt.Printf("1. %v\n", name)
 	flag.Parse()
 	fmt.Printf("2. %v\n", name)
 	fmt.Printf("Hello, %v!\n", name)
}
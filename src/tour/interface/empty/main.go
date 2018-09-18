package main

import "fmt"

func main() {
	var i interface{} // like `any` type in typescript
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

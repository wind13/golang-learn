
package main

import (
	"fmt"
  // "time"
)

var container = []string{"zero", "one", "two"}

func main() {
	aMap := map[string]int{
		"one": 1,
		"two": 2,
		"three": 3,
	}

	k := "two"
	v, ok := aMap[k]
	if ok {
		fmt.Printf("The element of key is %q: %d\n", k, v)
	} else {
		fmt.Println("Not found")
	}
	fmt.Println(uint8(200))
}

// go get -u -v github.com/acroca/go-symbols
// go get -u -v golang.org/x/tools/cmd/guru
// go get -u -v golang.org/x/tools/cmd/gorename
// go get -u -v github.com/derekparker/delve/cmd/dlv
// go get -u -v github.com/rogpeppe/godef
// go get -u -v github.com/sqs/goreturns
// go get -u -v github.com/golang/lint/golint
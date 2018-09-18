package main

import "fmt"
import "strings"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	fmt.Println(m)
	fmt.Println(m == nil)
	fmt.Println(m[""])
	fmt.Println(m["k"])
	fmt.Println(len(m))
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
	fmt.Println(len(m))
	elem, ok := m["Bell Labs"]
	elem0, ok0 := m["xyz"]
	fmt.Println(elem, ok)
	fmt.Println(elem0, ok0)
	fmt.Println(len(strings.Fields("That is not your falt.")))
}


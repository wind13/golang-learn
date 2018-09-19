package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	v := t.Value
	if nil != t.Left {
		Walk(t.Left, ch)
	}
	ch <- v
	if nil != t.Right {
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	Walk(t1, ch1)
	close(ch1)
	Walk(t2, ch2)
	close(ch2)
	for i := 0; i < 10; i++ {
		a := <-ch1
		b := <-ch2
		if a != b {
			return false
		}
	}
	return true
}

func main() {
	// ch := make(chan int, 10)
	t1 := tree.New(1)
	t2 := tree.New(1)
	// fmt.Println(t)
	// Walk(t, ch)
	// close(ch)
	// fmt.Println(len(ch))
	// for i := range ch {
	// fmt.Println(i)
	// }
	fmt.Println(Same(t1, t2))
}

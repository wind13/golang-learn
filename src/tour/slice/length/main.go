package main

import (
  "fmt"
  "strings"
)

// slice 向右扩展是随意的，不能向左扩展，但可以丢弃左边的元素。
func main() {
  var s []int
  fmt.Println(s, len(s), cap(s))
  if s == nil {
    fmt.Println("nil!")
  }

  s = []int{2, 3, 5, 7, 11, 13}
  printSlice("s", s)

  // Slice the slice to give it zero length.
  s = s[:0]
  printSlice("s", s)

  // Extend its length.
  s = s[:cap(s)]
  printSlice("s", s)

  // Drop its first two values.
  s = s[2:]
  printSlice("s", s)

  s = s[:4]
  printSlice("s", s)

  s = s[2:4]
  printSlice("s", s)

  a := make([]int, 5)
  printSlice("a", a)

  b := make([]int, 0, 5)
  printSlice("b", b)

  c := b[:2]
  printSlice("c", c)

  d := c[2:5]
  printSlice("d", d)

  e := make([]string, 4)
  printStrSlice("e", e)
  
  // Create a tic-tac-toe board.
  board := [][]string{
    []string{"_", "_", "_"},
    []string{"_", "_", "_"},
    []string{"_", "_", "_"},
  }

  // The players take turns.
  board[0][0] = "X"
  board[2][2] = "O"
  board[1][2] = "X"
  board[1][0] = "O"
  board[0][2] = "X"

  for i := 0; i < len(board); i++ {
    fmt.Printf("%s\n", strings.Join(board[i], " "))
  }

  var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
  for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
  }
  
  power := make([]int, 10)
	for i := range power {
		power[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range power {
		fmt.Printf("%d\n", value)
	}
}

// func printSlice(s []int) {
// 	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
// }

func printSlice(s string, x []int) {
  fmt.Printf("%s len=%d cap=%d %v\n",
    s, len(x), cap(x), x)
}

func printStrSlice(s string, x []string) {
  fmt.Printf("%s len=%d cap=%d %v\n",
    s, len(x), cap(x), x)
}
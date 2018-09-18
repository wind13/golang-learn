package main

import "fmt"

type T struct{
  i int
    d bool
}

func main() {
  var a [2]string
  a[0] = "Hello"
  a[1] = "World"
  fmt.Println(a[0], a[1])
  fmt.Println(a)

  primes := [6]int{2, 3, 5, 7, 11, 13}
  fmt.Println(primes)

  var s []int = primes[1:4]
  fmt.Println(s)

  names := [4]string{
    "John",
    "Paul",
    "George",
    "Ringo",
  }
  fmt.Println(names)

  c := names[0:2]
  b := names[1:3]
  fmt.Println(c, b)

  b[0] = "XXX"
  fmt.Println(c, b)
  fmt.Println(names)

  q := []int{2, 3, 5, 7, 11, 13}
  fmt.Println(q)

  r := []bool{true, false, true, true, false, true}
  fmt.Println(r)

  w := []struct{
    i int
    d bool
  }{
    {2, true},
    {3, false},
    {5, true},
    {7, true},
    {11, false},
    {13, true},
  }
  fmt.Println(w)
}

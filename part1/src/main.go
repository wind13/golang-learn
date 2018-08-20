package main

import (
  "flag"
  "fmt"
  "os"
  "libs"
)

var name string
var n float64

func init() {
  flag.StringVar(&name, "name", "everyone", "The greeting object.")
  flag.Float64Var(&n, "n", 9, "The sqrt number.")
}

func main() {
  flag.Usage = func() {
    fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
    flag.PrintDefaults()
  }
  flag.Parse()
  // fmt.Printf("Hello, %s!\nRandom:%s!\nTime:%s\n", name, rand.Intn(10), time.Now())
  libs.Hello(name)
  libs.Now()
  libs.Random()
  libs.Sqrt(n)
}
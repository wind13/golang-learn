package main

import (
  "flag"
  "fmt"
  "os"
  "part1/libs"
)
// import "./part1/libs/hello"

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
  fmt.Printf("Learning golang...\n")
  // Hello(name)
  libs.Hello(name)
  libs.Now()
  libs.Random()
  libs.Sqrt(n)
}
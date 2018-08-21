package main

import (
  "flag"
  "fmt"
  "os"
  "./libs"
  // "part1/libs"
)

// var name string
var n float64

func init() {
  // flag.StringVar(&name, "name", "everyone", "The greeting object.")
  flag.Float64Var(&n, "n", 9, "The sqrt number.")
}

func main() {
  // *type 是 cursor 指针 
  // 此处用 := 重定义了变量 name，不管在外面什么类型，都变为 *string 类型。
  // name := flag.String("name", "every one", "The greeting object.")

  var name *string
  name = flag.String("name", "every one", "The greeting object.")

  flag.Usage = func() {
    fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
    flag.PrintDefaults()
  }
  flag.Parse()
  fmt.Printf("Learning golang......\n")
  libs.Hello(*name)
  libs.Now()
  libs.Random()
  libs.Sqrt(n)
}
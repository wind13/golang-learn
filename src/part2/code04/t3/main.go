package main

import (
 "flag"
 "fmt"
)

func main() {
  // 此处 name 未指定类型，由编译时推断来确定，方便以后重构时此处不用修改。
  var name = getTheFlag()
  flag.Parse()
  fmt.Printf("Hello, %v!\n", *name)
}

func getTheFlag() *string {
  return flag.String("name", "everyone", "The greeting object.")
}
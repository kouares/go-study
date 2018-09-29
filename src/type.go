package main

import "fmt"

func main() {
  // 独自の名前で型を定義
  type hoge int
  type fuga string
  var h hoge = 123
  var f fuga = "test"
  fmt.Println(h)
  fmt.Println(f)
}

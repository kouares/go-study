package main

import "fmt"

func main() {
  // if
  var hoge = 100
  if hoge == 100 {
    fmt.Println("hoge = 100")
  } else if hoge < 100 {
    fmt.Println("hoge < 100")
  } else {
    fmt.Println("other case")
  }

  // ifで変数宣言
  // ifのスコープ内でのみ使用可能、elseでは使用できない
  if fuga := x(); fuga > 10 {
    fmt.Println(fuga)
  } else {
    fmt.Println(10)
  }
}

func x() int {
  return 10
}


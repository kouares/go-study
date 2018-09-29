package main

import (
  "fmt"
  "strconv"
)

func main() {
  // 定数宣言
  const hoge int = 100
  // 型を省略β
  const fuga = 100
  fmt.Println("hoge:" + strconv.Itoa(hoge))
  fmt.Println("fuga:" + strconv.Itoa(fuga))

  // 一括で定数を宣言
  const nemu, guru = 100, 200
  fmt.Println("nemu:" + strconv.Itoa(nemu))
  fmt.Println("guru:" + strconv.Itoa(guru))

  // 型が違っていてもOK
  const hoga, hogi = "test", 300
  fmt.Println("hoga:" + hoga)
  fmt.Println("hogi:" + strconv.Itoa(hogi))
}

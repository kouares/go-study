package main

import (
  "fmt"
  "strconv"
)

func main() {
  // 変数宣言
  var hoge int
  // 代入
  hoge = 100
  // 初期値
  var fuga int = 100
  // 型を省略
  var piyo = 100
  // varを省略
  kuru := 100
  fmt.Println("hoge:" + strconv.Itoa(hoge))
  fmt.Println("fuga:" + strconv.Itoa(fuga))
  fmt.Println("piyo:" + strconv.Itoa(piyo))
  fmt.Println("kuru:" + strconv.Itoa(kuru))

  // 一括で変数を宣言
  nemu, guru := 100, 200
  fmt.Println("nemu:" + strconv.Itoa(nemu))
  fmt.Println("guru:" + strconv.Itoa(guru))

  // 型が違っていてもOK
  hoga, hogi := "test", 300
  fmt.Println("hoga:" + hoga)
  fmt.Println("hogi:" + strconv.Itoa(hogi))
}

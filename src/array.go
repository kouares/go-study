package main

import "fmt"

func main() {
  // 配列宣言（空配列）
  var s1 = [3]string{}
  // 初期化あり
  var s2 = [3]string{"1", "2", "3"}
  // 要素数を省略
  var s3 = [...]string{"1", "2"}

  // 要素を取得、添字は0から
  fmt.Println(s1[0])

  // これはエラーになる
  //fmt.Println(s1[3])

  // 要素数取得
  fmt.Println("s2 length ", len(s2))
  fmt.Println("s3 length ", len(s3))
}

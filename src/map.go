package main

import "fmt"

func main() {
  // 普通のマップ
  // 末尾のコロンが必須なので注意
  hoge := map[string]int {
    "bob": 25,
    "alice": 24,
  }
  fmt.Println(hoge)

  // 要素取得
  fmt.Println(hoge["alice"])
  // 上書き
  hoge["alice"] = 27
  fmt.Println(hoge["alice"])
  // 追加
  hoge["leon"] = 35
  // 削除
  delete(hoge, "alice")
  fmt.Println(hoge)
}

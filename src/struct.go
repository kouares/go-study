package main

import "fmt"

func main() {
  // 構造体をuserという型で定義
  type user struct {
    Name string
    Age int
  }
  // データ指定なしで構造体を作成
  bob := user{}
  // 初期値あり（一部）
  bob2 := user {
    Name: "bob",
  }
  // 初期値あり（全部）
  bob3 := user {
    Name: "bob",
    Age: 25,
  }
  fmt.Println(bob)
  fmt.Println(bob2)
  fmt.Println(bob3)

  // 定義と同時に初期化する
  alice := struct {
    Name string
    Age int
  }{
    Name: "alice",
    Age: 24,
  }
  fmt.Println(alice)

  // 構造体のデータを変更
  alice.Name = "alice2"
  // 構造体のデータを取得
  fmt.Println(alice.Name)
}

package main

import "fmt"

func main() {
  //  ポインタ型の変数
  var p *int
  var n= 10

  // 変数nのアドレス（ポインタ型の値）をpに格納
  p = &n

  // pにはnのアドレス格納されているため、nを参照して値を取得
  fmt.Println(*p)

  // 関数とポインタの組み合わせ
  hoge := 1
  piyo := 1
  add1(hoge)
  add2(&piyo)
  fmt.Println(hoge, piyo)

  // 関数と構造体とポインタの組み合わせ
  bob := user{ Name: "bob" }
  alice := user{ Name: "alice" }
  rename1(bob)
  rename2(&alice)
  fmt.Println(bob)
  fmt.Println(alice)
}

// 渡された元のデータを参照しないデータを操作しているので値が変更されない
func add1(x int) { x = x + 1 }
// 渡された参照元の値を操作しているので値が変更される
func add2(y *int) { *y = *y + 1 }

type user struct {
  Name string
}

func rename1(x user) { x.Name = "leon" }
func rename2(y *user) { y.Name = "leon" }
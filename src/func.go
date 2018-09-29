package main

import "fmt"

func main() {
  fmt.Println("f1 result ", f1(1, 1))
  iresult, sresult := f2(2, "hoge")
  fmt.Println("f2 result ", iresult, sresult)

  // 関数を変数に代入して実行、いわゆる無名関数
  func1 := func () string { return "test" }
  func1()

  iresult, sresult = f3(2, "hoga")
  fmt.Println("f3 result ", iresult, sresult)
}

// 普通の関数
func f1(i1 int, i2 int) int {
  return i1 + i2
}

// 複数の戻り値を返す関数
func f2(i1 int, s1 string) (int , string) {
  return i1 + i1, s1
}


// 戻り値に名前を持たせた関数
func f3(i1 int, s1 string) (iresult int, sresult string) {
  iresult = i1 * i1 * i1
  sresult = s1 + s1
  return
}



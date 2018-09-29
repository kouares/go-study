package main

import "fmt"

func main() {
  // スライス宣言
  hoge := []string{"a", "b", "c", "d", "e"}
  fmt.Println(hoge)

  // 空のスライス
  // 配列は初期化をしないと、デフォルト値が設定されるが、スライスは空集合になる
  fuga := []string{}
  fmt.Println(fuga)

  // 要素取得
  // 配列からスライスを取得する
  // 省略した場合はそれぞれ先頭/末尾になる
  hoga := [5]string{"a", "b", "c", "d", "e"}
  s1 := hoga[1:2]
  s2 := hoga[1:]
  s3 := hoga[:2]
  s4 := hoga[:]
  fmt.Println(s1)
  fmt.Println(s2)
  fmt.Println(s3)
  fmt.Println(s4)

  // 要素の上書き
  // 元となった配列をスライスは参照する
  piyo := [5]string{"a", "b", "c", "d", "e"}
  piyo2 := piyo[:]
  fmt.Println(piyo)
  fmt.Println(piyo2)
  piyo[0] = "f"
  piyo2[1] = "g"
  fmt.Println(piyo)
  fmt.Println(piyo2)

  // 要素の追加
  fuga = append(fuga, "a")
  fmt.Println(fuga)

  // スライスのキャパシティ
  // スライスが保持できる要素数の上限、上限を超えた場合は元データへの参照が切られる
  // 上限が3の状態
  hogi := []string{"a", "b", "c"}
  hogi2 := hogi
  hogi2[0] = "d"
  fmt.Println(hogi)
  fmt.Println(hogi2)
  // 上限をこえるようにスライスに要素を追加する
  hogi = append(hogi, "d")
  fmt.Println(hogi)
  fmt.Println(hogi2)

  // キャパシティの確認
  // 初期サイズ3で要素数がキャパシティを上回る処理が行われると倍になる（倍のキャパシティをもった別のスライスが作成される）
  hogu := []string{"a", "b", "c"}
  fmt.Println("len", len(hogu))
  fmt.Println("cap", cap(hogu))
  hogu = append(hogu, "d")
  fmt.Println("len", len(hogu))
  fmt.Println("cap", cap(hogu))
}
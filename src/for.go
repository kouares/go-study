package main

import "fmt"

func main() {
  // for
  var sum int
  for i := 0; i < 10; i++ {
    sum += i
  }
  fmt.Println(sum)

  // for 無限ループ
  var loopSum int = 0
  for ;; {
    loopSum++
    if loopSum > 10 {
      break
    }
  }
  fmt.Println(loopSum)

  // whileはないのでforで
  var whileSum int = 0
  for whileSum < 10 {
    whileSum++
    fmt.Println("whileSum is", whileSum)
  }
}


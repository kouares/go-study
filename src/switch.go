package main

import "fmt"

func main() {
  // switch
  // seitchでpiyoを省略してcase式で評価することも出来る
  var piyo = 100
  switch piyo {
  case 0:
    fmt.Println("piyo = 0")
  case 100, 90, 80:
    fmt.Println("piyo = 100")
  default:
    fmt.Println("default")
  }

  // switch fallthrough
  var hoga = 100
  switch {
  case hoga > 0:
    fmt.Println("hoga > 0")
    fallthrough
  case hoga == 100:
    fmt.Println("hoga = 100")
    fallthrough
  default:
    fmt.Println("default")
  }
}

package main

import ( "fmt" )

func Sqrt(x float64) float64 {
  var y float64 = 1.0
  for i := 0; i < 10; i++ {
    y -= (y*y - x) / (2 * y)
  }
  return y
}

func main() {
  fmt.Println(Sqrt(2))
}

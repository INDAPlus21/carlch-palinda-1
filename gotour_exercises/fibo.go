package main

import ( "fmt" )

func fibonacci() func() int {
  n0, n1 := 0, 1
  return func() int {
    n0, n1 = n1, (n0 + n1)
    return (n1 - n0)
  }
}

func main() {
  f := fibonacci()
  for i := 0; i < 10; i++ {
    fmt.Println(f())
  }
}

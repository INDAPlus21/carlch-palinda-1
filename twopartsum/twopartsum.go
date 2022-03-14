package main

import ( "fmt" )

func sum(a []int, res chan<- int) {
  ret := 0
  for _, val := range a {
    ret += val
  }
  res <- ret
}

func ConcurrentSum(a []int) int {
  n := len(a)
  ch := make(chan int)
  go sum(a[:n/2], ch)
  go sum(a[n/2:], ch)

  return <-ch + <-ch
}

func main() {
  a := []int{1, 2, 3, 4, 5, 6, 7}
  fmt.Println(ConcurrentSum(a))
}

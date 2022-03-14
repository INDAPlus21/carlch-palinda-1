package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
  // Unfragmented 2d array
  matrix := make([][]uint8, dy)
  rows := make([]uint8, dy*dx)
  for i := range matrix {
    matrix[i] = rows[i*dx : (i+1)*dx]
    for j, _ := range matrix[i] {
      matrix[i][j] = uint8(i*i - j*j)
    }
  }
  return matrix
}

func main() {
  pic.Show(Pic)
}

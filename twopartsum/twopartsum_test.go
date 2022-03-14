package main

import ( "testing" )

func TestSumConcurrentCorrectlySumsEvenArray(t *testing.T) {
  arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
  expected := 55
  actual := ConcurrentSum(arr)
  if actual != expected {
    t.Errorf("expected %d, was %d", expected, actual)
  }
}

func TestSumConcurrentCorrectlySumsOddArray(t *testing.T) {
  arr := []int{1, 3, 5, 7, 9, 11, 13, 15}
  expected := 64
  actual := ConcurrentSum(arr)
  if actual != expected {
    t.Errorf("expected %d, was %d", expected, actual)
  }
}

func TestSumConcurrentCorrectlySumsNegativeArray(t *testing.T) {
  arr := []int{-1, -2, -3, -4, -5, -6, -7, -8, -9, -10}
  expected := -55
  actual := ConcurrentSum(arr)
  if actual != expected {
    t.Errorf("expected %d, was %d", expected, actual)
  }
}

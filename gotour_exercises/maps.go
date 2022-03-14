package main

import(
  "strings"

  "golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
  word := strings.Fields(s)
  count := make(map[string]int)
  for i := range word {
    count[word[i]]++
  }
  return count
}

func main() {
  wc.Test(WordCount)
}

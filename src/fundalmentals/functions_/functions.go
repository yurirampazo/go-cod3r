package main

import "strconv"


func sum(a int, b int) int {
  return a + b
}

func print2(value string) {
  println(value)
}


func main() {
  result := sum(1,3)
  print2(strconv.Itoa(result))
}
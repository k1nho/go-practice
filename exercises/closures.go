package main

import "fmt"


func fibonacci() func() int {
  prev := 0
  sum := 1

  return func() int {
    next := prev + sum
    prev = sum
    sum = next
    return sum
  }
}

func main() {
  f := fibonacci()

  for i := 0; i < 10; i++ {
    fmt.Printf("fibonacci of %d\n", i+1)
    fmt.Println(f())
  }
}

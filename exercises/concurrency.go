package main

import (
	"fmt"
	"sync"
	"time"
)


func printMessage(wg *sync.WaitGroup) {
  fmt.Println("here is the message")
  wg.Done()
}

func main() {
  t := time.Now()
  connections := 3

  var wg sync.WaitGroup
  wg.Add(connections)

  for i := 0; i < connections; i++ {
    go printMessage(&wg)
  }
  
  wg.Wait()
  
  fmt.Printf("time passed %v\n", time.Since(t))
  fmt.Println("main finish")
}


package main

import (
	"fmt"
	"runtime"
)

// switch
func main() {
  fmt.Print("Running go on ")
  switch os := runtime.GOOS; os {
  case "darwin": 
    fmt.Println("OS X")
  case "linux" :
    fmt.Println("Penguin power")
  default:
  fmt.Printf("running %s\n", os);

  }
}

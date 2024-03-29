package main

import (
	"fmt"
	"sync"
	"time"
)


type Tree struct {
  Left *Tree
  Right *Tree
  Value int
}

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
  
  fmt.Println("for select pattern")
  forSelectPattern()

  ch := make(chan int)
  root := Tree{
    Left: &Tree{nil, nil, 0},
    Right: &Tree{nil, nil, 2},
    Value : 1,
  }

  go func(){
    Walk(&root, ch)
    close(ch)
  }()

  for i := range ch {
    fmt.Printf("%d ", i)
  }

  
  fmt.Printf("time passed %v\n", time.Since(t))
  fmt.Println("Finish")
}

func Walk(t *Tree, ch chan int) {
  if t == nil {
    return
  }
  Walk(t.Left, ch)
  ch <- t.Value
  Walk(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *Tree) bool {
	cht1 := make(chan int)
	cht2 := make(chan int)
	
	go func() {
		Walk(t1, cht1)
		close(cht1)
	}()
	
	go func(){
		Walk(t2, cht2)
		close(cht2)
	}()
	
	for {
		v1, ok1 := <-cht1
		v2, ok2 := <-cht2
		
		if ok1 != ok2 || v1 != v2 {
			return false
		}
		
		if !ok1 && !ok2 {
			break
		}
	}
	
	return true
	
}


// patterns

// 1. for-select loop (buffered channels)
// make go routine queue up results in a channel
func forSelectPattern() {
  channel := make(chan string, 3)
  leaders := []string{"Brock", "Blaine", "Cynthia"}

  for _, leader := range leaders {
    select {
      case channel <- leader:
   }
  }

  close(channel)

  for data := range channel {
    fmt.Printf("Leader obtained %s\n", data)
  }
  
}



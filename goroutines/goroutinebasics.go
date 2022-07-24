package main

import (
	"fmt"
	"sync"
)

func push(name string, ch chan string) {
	msg := "Hey " + name
	ch <- msg
}

func main() {
	//creating a channel
	ch := make(chan string)

	// Starting concurrent routines
	go push("Kin", ch)
	go push("Subaru", ch)
	go push("Kazuma", ch)

	// Order is not guaranteed since they are concurrent
	fmt.Println(<-ch, <-ch, <-ch)

	integerSub()
	runWaitGroup()
}

func integerSub() {
	// specify limit for the amount of messages it can keep
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	// adding a 3rd subscription here will create a deadlock since the channel is limited to 2 concurrent connections at max
	// ch <- 3

	// We can see the status of the routine (true if it is open) and v is the value
	v, ok := <-ch
	fmt.Println(v, ok)
	// To close the the channel we can use close
	close(ch)
}

// We can use a wait group to have a collection of goroutines to finish
// we can use Add to set the number of goroutines to wait for
// we can then call done when it finishes
func runWaitGroup() {
	itemList := []string{"Steak", "Fish", "Chicken", "Cup"}
	var wg sync.WaitGroup

	for _, item := range itemList {
		wg.Add(1)
		go operate(&wg, item)
	}
	wg.Wait()
}

func operate(wg *sync.WaitGroup, item string) {
	defer wg.Done()
	item = item + " this is item"
	fmt.Println(item)
}

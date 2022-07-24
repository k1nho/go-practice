package main

import (
	"fmt"
	"sync"
	"time"
)

// structs
type Point struct {
	X int
	Y int
}

// Receivers (getters)
func (p Point) Diff() int {
	res := p.X - p.Y
	return res
}

// mutations (setters)
func (p *Point) Scale(factor int) {
	p.X = p.X * factor
	p.Y = p.Y * factor
}

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
	deferrer()

	fmt.Println("Defining  a point")
	p := Point{X: 2, Y: 2}
	fmt.Printf("The x coordinate is %v and the y coordinate is %v\n", p.X, p.Y)
	p.X = 5
	fmt.Printf("Changing the code to assign x to a %v\n", p.X)
	// lets use getters to get the difference
	diff := p.Diff()
	fmt.Printf("Difference between x and y is %v\n", diff)
	p.Scale(2)
	fmt.Printf("Doubling our coordinates we have x: %v, y: %v\n", p.X, p.Y)

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

// Defers a function running until the surrounding function returns
func deferrer() {
	var d = int64(0)
	defer func(d *int64) {
		fmt.Printf(" & %v seconds", *d)
	}(&d)
	fmt.Print("Done\n")
	d = time.Now().Unix()
}

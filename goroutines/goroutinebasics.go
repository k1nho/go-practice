package main

import "fmt"

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
}

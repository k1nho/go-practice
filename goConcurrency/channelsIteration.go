package main

import (
	"fmt"
)

func main() {
	channel := make(chan string)
	n := 3

	// start concurrent routine
	go shootBall(channel, n)

	for i := 0; i < n; i++ {
		fmt.Println(<-channel)
	}
}

func shootBall(channel chan string, rounds int) {
	for i := 0; i < rounds; i++ {
		channel <- fmt.Sprint("Shoot ball!")
	}
}

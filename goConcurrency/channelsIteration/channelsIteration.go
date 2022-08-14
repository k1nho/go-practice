package main

import (
	"fmt"
)

func main() {
	channel := make(chan string)
	channel2 := make(chan string)
	n := 3

	// start concurrent routine
	go shootBall(channel, n)

	for i := 0; i < n; i++ {
		fmt.Println(<-channel)
	}

	go throwBall(channel2, n)
	for message := range channel2 {
		fmt.Println(message)
	}
}

func shootBall(channel chan string, rounds int) {
	for i := 0; i < rounds; i++ {
		channel <- fmt.Sprint("Shoot ball!")
	}
}

func throwBall(channel chan string, rounds int) {
	for i := 0; i < rounds; i++ {
		channel <- fmt.Sprint("Throw ball!")
	}

	// using close method to signal to the channel that we have finished our processes
	close(channel)
}

package main

import (
	"fmt"
	"time"
)

func main() {

	// lets calculate the time it takes to get all medals
	start := time.Now()

	defer func() {
		fmt.Println(time.Since(start))
	}()

	// we need a way that is better to check when all our routines finish so that we can finish the main function at that time
	// One trivial way to do it is to have the main function sleep for some seconds (time.Sleep(time.Second*2)) which allow the goroutines to be finished however
	// we will encounter that this is not very good because the idea is that we want to save time and eventhough we slept our main we still
	// are wasting time. Thus, we will use channels to create a shared mem so that we can now when all our processes finish
	medalsCompleted := make(chan int)

	gymLeaders := []string{"Brock", "Erika", "Blaine", "Giovanni"}

	for _, leader := range gymLeaders {
		// we can either do it like the following: getMedal(leader) in which case it will take us 4 seconds
		// or we can make it concurrent using goroutines
		go getMedal(leader, medalsCompleted)
	}

	fmt.Println(<-medalsCompleted) // now the completion time is closer to what it actually takes to complete our routines in parallel that is 1 sec
}

func getMedal(gymLeader string, medalsCompleted chan int) {
	fmt.Printf("Get medal from %v\n", gymLeader)
	time.Sleep(time.Second)
	medalsCompleted <- 1
}

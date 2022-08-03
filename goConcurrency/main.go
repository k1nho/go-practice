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

	gymLeaders := []string{"Brock", "Erika", "Blaine", "Giovanni"}

	for _, leader := range gymLeaders {
		// we can either do it like the following: getMedal(leader) in which case it will take us 4 seconds
		// or we can make it concurrent using goroutines
		go getMedal(leader)
	}

	time.Sleep(time.Second * 2)
}

func getMedal(gymLeader string) {
	fmt.Printf("Get medal from %v\n", gymLeader)
	time.Sleep(time.Second)

}

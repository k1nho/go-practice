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
		getMedal(leader)
	}
}

func getMedal(gymLeader string) {
	fmt.Printf("Get medal from %v\n", gymLeader)
	time.Sleep(time.Second)

}

package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	count int
	lock  sync.Mutex
)

func main() {
	numberOfRoutines := 10000

	for i := 0; i < numberOfRoutines; i++ {
		go increaseCount()
	}

	time.Sleep(time.Second * 1)
	fmt.Printf("The count is %v", count)
}

func increaseCount() {
	lock.Lock()
	count++ // this operation is not atomic, since it actually performs a reassignement to get the correct result
	lock.Unlock()
}

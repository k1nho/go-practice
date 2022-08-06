// we can use atomic for some cases instead of mutex

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var (
	count int32
)

func main() {
	// setup
	workers := 100000000

	var wg sync.WaitGroup
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go increaseCount(&wg)
	}
	wg.Wait()
	fmt.Printf("The value of count is %d", count)
}

func increaseCount(wg *sync.WaitGroup) {
	atomic.AddInt32(&count, 1)
	wg.Done()
}

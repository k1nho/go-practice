package main

import (
	"fmt"
	"sync"
)

func main() {
	var memory int

	// create a resource pool where we can either create new resources if there are not available resources or have the resources be emplaced back in there
	memoryPool := &sync.Pool{
		New: func() interface{} {
			memory++
			mem := make([]byte, 1024)
			return &mem
		},
	}

	workers := 1024 * 1024 //worker threads that will do some work
	var wg sync.WaitGroup
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func() {
			resource := memoryPool.Get().(*[]byte)
			// we do some work here
			fmt.Sprintln("Doing some work...")
			// return resource to the pool
			memoryPool.Put(resource)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("Process has finished! total number of resources created: %d", memory)
}

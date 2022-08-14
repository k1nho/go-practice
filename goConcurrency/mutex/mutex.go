package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	count  int
	lock   sync.Mutex
	rwlock sync.RWMutex
)

func main() {
	numberOfRoutines := 10000
	numberOfReaders := 5
	for i := 0; i < numberOfRoutines; i++ {
		go increaseCount()
	}

	go writer()
	// we can also use a read/write lock to be able to access the critical section by many readers but only one writer
	for i := 0; i < numberOfReaders; i++ {
		go reader()
	}

	time.Sleep(time.Second * 3)
	fmt.Printf("The count is %v", count)
}

func increaseCount() {
	lock.Lock()
	count++ // this operation is not atomic, since it actually performs a reassignement to get the correct result
	lock.Unlock()
}

func reader() {
	rwlock.RLock()
	defer rwlock.RUnlock()
	fmt.Println("Reader acquire lock")
	time.Sleep(time.Second * 1)
	fmt.Println("Reader release lock")
}

func writer() {
	rwlock.Lock()
	defer rwlock.Unlock()
	fmt.Println("Writer acquire lock")
	time.Sleep(time.Second * 1)
	fmt.Println("Writer release lock")
}

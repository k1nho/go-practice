package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var assertReady bool

func main() {
	var workTime int
	cond := sync.NewCond(&sync.Mutex{})

	go prepareQuery(cond)

	cond.L.Lock()
	if !assertReady {
		workTime++
		cond.Wait()
	}
	cond.L.Unlock()

	broadcastQuery()
	fmt.Printf("The total work intervals were %d\n", workTime)
}

func prepareQuery(cond *sync.Cond) {
	sleep()
	assertReady = true
	cond.Signal()
}

func sleep() {
	rand.Seed(time.Now().UnixNano())
	waitTime := time.Duration(1+rand.Intn(5)) * time.Second
	time.Sleep(waitTime)
}

func broadcastQuery() {
	ready := sync.NewCond(&sync.Mutex{})
	var wg sync.WaitGroup
	wg.Add(3)

	for i := 0; i < 3; i++ {
		waitForGroupQuery(spinQuery, ready, i, &wg)
	}
	ready.Broadcast()
	wg.Wait()
	fmt.Println("Queries have been started")
}

func spinQuery(queryNum int, wg *sync.WaitGroup) {
	fmt.Printf("Query %d has started\n", queryNum)
	wg.Done()
}

// Using Broadcast
func waitForGroupQuery(fn func(int, *sync.WaitGroup), ready *sync.Cond, i int, waitGroup *sync.WaitGroup) {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		wg.Done()
		ready.L.Lock()
		defer ready.L.Unlock()
		ready.Wait()
		fn(i, waitGroup)
	}()

	wg.Wait()
}

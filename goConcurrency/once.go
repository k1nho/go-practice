package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	querySuccessful bool
)

func main() {
	// In this example we will use Once in sync package to only execute a function once
	// Lets say we want to make sure a query is successful; however, it is hard for it to be sucessful so we will spawn some routines to try multiple times
	n := 1000
	var wg sync.WaitGroup
	var once sync.Once
	wg.Add(n)
	for i := 0; i < n; i++ {
		if retryQuery() {
			once.Do(markQuerySucess)
		}
		wg.Done()
	}

	wg.Wait()
	fmt.Printf("Status of the query %v", querySuccessful)
}

func markQuerySucess() {
	fmt.Println("Query has been successful!")
	querySuccessful = true
}

func retryQuery() bool {
	rand.Seed(time.Now().UnixNano())
	return 1 == rand.Intn(10)
}

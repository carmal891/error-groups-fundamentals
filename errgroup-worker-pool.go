package main

import (
	"fmt"
	"sync"
	"time"
)

// Data to be proccessed.
var taskCount = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

const (
	// Number of concurrent workers.
	numberOfWorkers = 3
)

func eGrpWorkerPool() {
	// Create buffered channel.
	jobs := make(chan struct{}, numberOfWorkers)
	wg := sync.WaitGroup{}

	// Add workers.
	for id := range taskCount {
		wg.Add(1)
		jobs <- struct{}{}

		go func(id int) {
			worker(id)
			<-jobs
			defer wg.Done()
		}(id)
	}

	// Wait for all workers to complete.
	wg.Wait()
}

func worker(id int) {
	fmt.Println(id)
	time.Sleep(2 * time.Second)
}

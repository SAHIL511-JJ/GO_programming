// Program for checkpoint synchronization
package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup, barrier *sync.WaitGroup) {
	defer wg.Done()

	// simulating work
	fmt.Printf("Worker %d: starting work\n", id)
	time.Sleep(time.Duration(id*100) * time.Millisecond)
	fmt.Printf("Worker %d: finished part, waiting at checkpoint\n", id)

	// signal that this worker reached checkpoint
	barrier.Done()

	// wait for all workers at checkpoint
	barrier.Wait()

	// after checkpoint, continue with assembly
	fmt.Printf("Worker %d: putting parts together\n", id)
}

func main() {
	numWorkers := 4

	var wg sync.WaitGroup
	var barrier sync.WaitGroup

	barrier.Add(numWorkers)

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, &wg, &barrier)
	}

	wg.Wait()
	fmt.Println("All workers completed assembly!")
}

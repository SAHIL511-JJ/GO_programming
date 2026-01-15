// Program with 5 goroutines generating numbers 0-10
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func generateNumbers(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i <= 10; i++ {
		fmt.Printf("Goroutine %d: %d\n", id, i)
		// random delay between 0-250 ms
		delay := rand.Intn(250)
		time.Sleep(time.Duration(delay) * time.Millisecond)
	}
	fmt.Printf("Goroutine %d finished\n", id)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var wg sync.WaitGroup

	// launching 5 goroutines
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go generateNumbers(i, &wg)
	}

	// waiting for all goroutines to finish
	wg.Wait()
	fmt.Println("All goroutines completed!")
}

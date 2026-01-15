// Program to check even/odd using channels
package main

import (
	"fmt"
	"sync"
)

func evenReceiver(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range ch {
		fmt.Println("Even goroutine received:", num)
	}
}

func oddReceiver(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range ch {
		fmt.Println("Odd goroutine received:", num)
	}
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	evenCh := make(chan int)
	oddCh := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2)

	// starting receiver goroutines
	go evenReceiver(evenCh, &wg)
	go oddReceiver(oddCh, &wg)

	// sending numbers to respective channels
	for _, num := range numbers {
		if num%2 == 0 {
			evenCh <- num
		} else {
			oddCh <- num
		}
	}

	// closing channels
	close(evenCh)
	close(oddCh)

	wg.Wait()
	fmt.Println("Done!")
}

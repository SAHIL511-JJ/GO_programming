// Program to demonstrate closing channel with for-range
package main

import "fmt"

func producer(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i
		fmt.Println("Sent:", i)
	}
	// closing channel after sending all values
	close(ch)
	fmt.Println("Channel closed")
}

func main() {
	ch := make(chan int)

	// starting producer goroutine
	go producer(ch)

	// reading using for-range (automatically stops when channel closed)
	fmt.Println("\nReceiving values:")
	for value := range ch {
		fmt.Println("Received:", value)
	}

	fmt.Println("All values received!")
}

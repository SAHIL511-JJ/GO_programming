// Program to read/write Fibonacci to channel
package main

import "fmt"

func fibonacci(n int, ch chan int) {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		ch <- a
		a, b = b, a+b
	}
	close(ch)
}

func main() {
	ch := make(chan int, 10)

	// writing fibonacci to channel
	go fibonacci(10, ch)

	// reading from channel
	fmt.Println("Fibonacci Series:")
	for num := range ch {
		fmt.Print(num, " ")
	}
	fmt.Println()
}

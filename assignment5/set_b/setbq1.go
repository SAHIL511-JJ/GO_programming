// Program to demonstrate buffered channel
package main

import "fmt"

func main() {
	// creating buffered channel with capacity 5
	ch := make(chan string, 5)

	// storing values
	ch <- "Apple"
	ch <- "Banana"
	ch <- "Orange"

	fmt.Println("Channel capacity:", cap(ch))
	fmt.Println("Channel length:", len(ch))

	// reading values
	fmt.Println("\nReading from channel:")
	fmt.Println(<-ch)
	fmt.Println("Length after reading:", len(ch))

	fmt.Println(<-ch)
	fmt.Println("Length after reading:", len(ch))

	fmt.Println(<-ch)
	fmt.Println("Length after reading:", len(ch))
}

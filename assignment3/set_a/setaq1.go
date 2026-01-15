// Program to find largest and smallest in array
package main

import "fmt"

func main() {
	arr := []int{45, 12, 78, 23, 56, 89, 34}

	fmt.Println("Array:", arr)

	// assuming first element is both largest and smallest
	largest := arr[0]
	smallest := arr[0]

	for i := 0; i < len(arr); i++ {
		if arr[i] > largest {
			largest = arr[i]
		}
		if arr[i] < smallest {
			smallest = arr[i]
		}
	}

	fmt.Println("Largest:", largest)
	fmt.Println("Smallest:", smallest)
}

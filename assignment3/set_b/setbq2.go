// Program to sort array in ascending order
package main

import "fmt"

func main() {
	arr := []int{64, 25, 12, 22, 11}

	fmt.Println("Before sorting:", arr)

	// bubble sort algorithm
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				// swap
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	fmt.Println("After sorting:", arr)
}

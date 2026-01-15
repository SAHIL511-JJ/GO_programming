// Program to demonstrate slice operations
package main

import "fmt"

func main() {
	// creating a slice
	slice := []int{10, 20, 30, 40, 50}
	fmt.Println("Original slice:", slice)
	fmt.Println("Length:", len(slice), "Capacity:", cap(slice))

	// append operation
	slice = append(slice, 60, 70)
	fmt.Println("\nAfter append(60, 70):", slice)

	// copy operation
	slice2 := make([]int, len(slice))
	copy(slice2, slice)
	fmt.Println("Copied slice:", slice2)

	// slicing operation
	subSlice := slice[2:5]
	fmt.Println("\nSub-slice [2:5]:", subSlice)

	// remove element at index 2
	index := 2
	slice = append(slice[:index], slice[index+1:]...)
	fmt.Println("After removing element at index 2:", slice)

	// inserting element at index 2
	element := 100
	slice = append(slice[:index], append([]int{element}, slice[index:]...)...)
	fmt.Println("After inserting 100 at index 2:", slice)

	fmt.Println("\nFinal Length:", len(slice))
	fmt.Println("Final Capacity:", cap(slice))
}

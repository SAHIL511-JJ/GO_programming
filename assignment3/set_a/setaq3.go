// Program to initialize slice using multi-line syntax
package main

import "fmt"

func main() {
	// slice with multi-line syntax
	fruits := []string{
		"Apple",
		"Banana",
		"Orange",
		"Mango",
		"Grapes",
	}

	fmt.Println("Slice of Fruits:")
	fmt.Println(fruits)

	fmt.Println("\nDisplaying each element:")
	for i := 0; i < len(fruits); i++ {
		fmt.Println(i+1, ":", fruits[i])
	}

	fmt.Println("\nLength of slice:", len(fruits))
	fmt.Println("Capacity of slice:", cap(fruits))
}

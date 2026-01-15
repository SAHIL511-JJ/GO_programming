// Program to copy array elements using method
package main

import "fmt"

// Array type
type IntArray []int

// method to copy array
func (a IntArray) CopyTo(dest IntArray) {
	for i := 0; i < len(a) && i < len(dest); i++ {
		dest[i] = a[i]
	}
}

func main() {
	source := IntArray{10, 20, 30, 40, 50}
	dest := make(IntArray, len(source))

	fmt.Println("Source array:", source)
	fmt.Println("Destination before copy:", dest)

	// copying using method
	source.CopyTo(dest)

	fmt.Println("Destination after copy:", dest)
}

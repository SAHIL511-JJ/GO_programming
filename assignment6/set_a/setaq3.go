// Program to find area of rectangle
package main

import "fmt"

// Area function for rectangle
func Area(length, width float64) float64 {
	return length * width
}

func main() {
	var length, width float64

	fmt.Print("Enter length of rectangle: ")
	fmt.Scan(&length)

	fmt.Print("Enter width of rectangle: ")
	fmt.Scan(&width)

	area := Area(length, width)
	fmt.Printf("Area of rectangle: %.2f\n", area)
}

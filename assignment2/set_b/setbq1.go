// Program to swap using call by reference
package main

import "fmt"

// function that swaps using pointers (call by reference)
func swap(x *int, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}

func main() {
	var a, b int

	fmt.Print("Enter first number: ")
	fmt.Scan(&a)

	fmt.Print("Enter second number: ")
	fmt.Scan(&b)

	fmt.Println("Before swap: a =", a, "b =", b)

	// passing addresses of a and b
	swap(&a, &b)

	fmt.Println("After swap: a =", a, "b =", b)
}

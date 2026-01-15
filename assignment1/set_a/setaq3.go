// Program to swap two numbers without using temporary variable
package main

import "fmt"

func main() {
	var a, b int

	fmt.Print("Enter first number: ")
	fmt.Scan(&a)
	fmt.Print("Enter second number: ")
	fmt.Scan(&b)

	fmt.Println("Before swap:")
	fmt.Println("a =", a, "b =", b)

	// swapping without temp variable
	a = a + b
	b = a - b
	a = a - b

	fmt.Println("After swap:")
	fmt.Println("a =", a, "b =", b)
}

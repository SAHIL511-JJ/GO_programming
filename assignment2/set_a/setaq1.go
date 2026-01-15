// Program to add two numbers using function
package main

import "fmt"

// function to add two numbers
func add(a int, b int) int {
	var result int
	result = a + b
	return result
}

func main() {
	var num1, num2 int

	fmt.Print("Enter first number: ")
	fmt.Scan(&num1)

	fmt.Print("Enter second number: ")
	fmt.Scan(&num2)

	sum := add(num1, num2)
	fmt.Println("Sum =", sum)
}

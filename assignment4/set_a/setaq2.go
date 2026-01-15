// Program to multiply two numbers using method
package main

import "fmt"

// Number type
type Number int

// multiply method
func (n Number) Multiply(m Number) Number {
	return n * m
}

func main() {
	var num1 Number = 5
	var num2 Number = 7

	result := num1.Multiply(num2)

	fmt.Println("Number 1:", num1)
	fmt.Println("Number 2:", num2)
	fmt.Println("Multiplication:", result)
}

// Simple calculator program
package main

import "fmt"

// calculator functions
func add(a, b float64) float64      { return a + b }
func subtract(a, b float64) float64 { return a - b }
func multiply(a, b float64) float64 { return a * b }
func divide(a, b float64) float64   { return a / b }

func main() {
	var num1, num2 float64
	var choice int

	fmt.Println("Calculator")
	fmt.Println("1. Addition")
	fmt.Println("2. Subtraction")
	fmt.Println("3. Multiplication")
	fmt.Println("4. Division")

	fmt.Print("Enter choice: ")
	fmt.Scan(&choice)

	fmt.Print("Enter first number: ")
	fmt.Scan(&num1)
	fmt.Print("Enter second number: ")
	fmt.Scan(&num2)

	var result float64
	switch choice {
	case 1:
		result = add(num1, num2)
		fmt.Println("Result:", result)
	case 2:
		result = subtract(num1, num2)
		fmt.Println("Result:", result)
	case 3:
		result = multiply(num1, num2)
		fmt.Println("Result:", result)
	case 4:
		if num2 != 0 {
			result = divide(num1, num2)
			fmt.Println("Result:", result)
		} else {
			fmt.Println("Error: Division by zero")
		}
	default:
		fmt.Println("Invalid choice")
	}
}

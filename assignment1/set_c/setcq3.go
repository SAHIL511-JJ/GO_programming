// Simple calculator using switch case
package main

import "fmt"

func main() {
	var num1, num2 float64
	var choice int

	fmt.Println("Simple Calculator")
	fmt.Println("-----------------")
	fmt.Println("1. Addition")
	fmt.Println("2. Subtraction")
	fmt.Println("3. Multiplication")
	fmt.Println("4. Division")

	fmt.Print("Enter your choice (1-4): ")
	fmt.Scan(&choice)

	fmt.Print("Enter first number: ")
	fmt.Scan(&num1)

	fmt.Print("Enter second number: ")
	fmt.Scan(&num2)

	switch choice {
	case 1:
		fmt.Println("Result:", num1+num2)
	case 2:
		fmt.Println("Result:", num1-num2)
	case 3:
		fmt.Println("Result:", num1*num2)
	case 4:
		if num2 != 0 {
			fmt.Println("Result:", num1/num2)
		} else {
			fmt.Println("Cannot divide by zero!")
		}
	default:
		fmt.Println("Invalid choice!")
	}
}

// Program to find recursive sum of digits
package main

import "fmt"

// recursive function to find sum of digits
func sumOfDigits(n int) int {
	if n == 0 {
		return 0
	}
	return (n % 10) + sumOfDigits(n/10)
}

func main() {
	var num int

	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	// handling negative numbers
	if num < 0 {
		num = -num
	}

	result := sumOfDigits(num)
	fmt.Println("Sum of digits =", result)
}

// Program to print Fibonacci series
package main

import "fmt"

func main() {
	var n int

	fmt.Print("Enter how many terms: ")
	fmt.Scan(&n)

	var first int = 0
	var second int = 1

	fmt.Println("Fibonacci Series:")

	for i := 1; i <= n; i++ {
		fmt.Print(first, " ")

		// calculating next term
		next := first + second
		first = second
		second = next
	}
	fmt.Println()
}

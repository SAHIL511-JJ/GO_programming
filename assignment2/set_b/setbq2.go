// Program to demonstrate named return variables
package main

import "fmt"

// function with named return variables
func calculate(num int) (square int, cube int) {
	// we can directly use square and cube
	square = num * num
	cube = num * num * num
	return // no need to mention what to return
}

func main() {
	var n int

	fmt.Print("Enter a number: ")
	fmt.Scan(&n)

	sq, cu := calculate(n)

	fmt.Println("Square of", n, "=", sq)
	fmt.Println("Cube of", n, "=", cu)
}

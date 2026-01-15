// Program to illustrate call by value
package main

import "fmt"

// function that tries to change value (call by value)
func changeValue(x int) {
	x = x + 100
	fmt.Println("Inside function x =", x)
}

func main() {
	var num int = 50

	fmt.Println("Before function call: num =", num)

	// passing value (not address)
	changeValue(num)

	// original value remains unchanged
	fmt.Println("After function call: num =", num)

	// This shows that in call by value
	// changes inside function dont affect original variable
}

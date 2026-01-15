// Program to explain new function in Go
package main

import "fmt"

func main() {
	// using new() to create pointer
	numPtr := new(int)

	fmt.Println("Value at numPtr:", *numPtr) // default value is 0
	fmt.Println("Address stored:", numPtr)

	// assigning value using pointer
	*numPtr = 42

	fmt.Println("After assigning 42:")
	fmt.Println("Value at numPtr:", *numPtr)

	// another example with string
	strPtr := new(string)
	fmt.Println()
	fmt.Println("String pointer default:", *strPtr) // empty string
	*strPtr = "Hello Go"
	fmt.Println("After assigning:", *strPtr)
}

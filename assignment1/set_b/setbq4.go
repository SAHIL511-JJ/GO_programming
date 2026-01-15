// Program to demonstrate pointer to pointer
package main

import "fmt"

func main() {
	var num int = 100

	// pointer to num
	var ptr *int = &num

	// pointer to pointer
	var pptr **int = &ptr

	fmt.Println("Value of num:", num)
	fmt.Println("Address of num:", &num)
	fmt.Println()

	fmt.Println("ptr holds address:", ptr)
	fmt.Println("Value at ptr (*ptr):", *ptr)
	fmt.Println("Address of ptr:", &ptr)
	fmt.Println()

	fmt.Println("pptr holds address:", pptr)
	fmt.Println("Value at pptr (*pptr):", *pptr)
	fmt.Println("Value at **pptr:", **pptr)
}

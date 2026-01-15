// Program to print address of a variable
package main

import "fmt"

func main() {
	var num int = 50
	var name string = "GoLang"

	fmt.Println("Value of num:", num)
	fmt.Println("Address of num:", &num)

	fmt.Println()

	fmt.Println("Value of name:", name)
	fmt.Println("Address of name:", &name)
}

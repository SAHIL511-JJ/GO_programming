// Program to demonstrate type assertion
package main

import "fmt"

func main() {
	// creating interface variable
	var i interface{} = "Hello World"

	// type assertion
	s := i.(string)
	fmt.Println("Value:", s)

	// safe type assertion with ok
	s2, ok := i.(string)
	fmt.Println("Value:", s2, "OK:", ok)

	// checking for wrong type
	n, ok := i.(int)
	fmt.Println("Int value:", n, "OK:", ok)

	// another example
	var value interface{} = 100
	if v, ok := value.(int); ok {
		fmt.Println("Integer value:", v)
	}
}

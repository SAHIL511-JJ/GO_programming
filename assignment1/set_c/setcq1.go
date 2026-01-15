// Program to concatenate two strings using pointers
package main

import "fmt"

func main() {
	var str1 string = "Hello "
	var str2 string = "World"

	// creating pointers to strings
	ptr1 := &str1
	ptr2 := &str2

	fmt.Println("String 1:", *ptr1)
	fmt.Println("String 2:", *ptr2)

	// concatenating using pointers
	result := *ptr1 + *ptr2

	fmt.Println("Concatenated string:", result)
}

// Program to compare two strings
package main

import "fmt"

func main() {
	var str1, str2 string

	fmt.Print("Enter first string: ")
	fmt.Scan(&str1)

	fmt.Print("Enter second string: ")
	fmt.Scan(&str2)

	// comparing strings
	if str1 == str2 {
		fmt.Println("Both strings are equal")
	} else if str1 > str2 {
		fmt.Println(str1, "is greater than", str2)
	} else {
		fmt.Println(str1, "is less than", str2)
	}
}

// Program to check if first string is substring of second
package main

import (
	"fmt"
	"strings"
)

func main() {
	var str1, str2 string

	fmt.Print("Enter first string: ")
	fmt.Scan(&str1)

	fmt.Print("Enter second string: ")
	fmt.Scan(&str2)

	// checking if str1 is substring of str2
	if strings.Contains(str2, str1) {
		fmt.Println(str1, "is a substring of", str2)
	} else {
		fmt.Println(str1, "is NOT a substring of", str2)
	}
}

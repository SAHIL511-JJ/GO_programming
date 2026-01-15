// Program to check if number is single digit
package main

import "fmt"

func main() {
	var num int

	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	// single digit means between -9 to 9
	if num >= -9 && num <= 9 {
		fmt.Println(num, "is a single digit number")
	} else {
		fmt.Println(num, "is not a single digit number")
	}
}

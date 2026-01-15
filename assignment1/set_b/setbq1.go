// Program to print multiplication table of a number
package main

import "fmt"

func main() {
	var num int

	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	fmt.Println("Table of", num)
	fmt.Println("--------------")

	for i := 1; i <= 10; i++ {
		result := num * i
		fmt.Println(num, "x", i, "=", result)
	}
}

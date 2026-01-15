// Program to print Pascal's Triangle
package main

import "fmt"

func main() {
	var rows int

	fmt.Print("Enter number of rows: ")
	fmt.Scan(&rows)

	fmt.Println("Pascal's Triangle")
	fmt.Println("-----------------")

	for i := 0; i < rows; i++ {
		// printing spaces for alignment
		for j := 0; j < rows-i-1; j++ {
			fmt.Print(" ")
		}

		var num int = 1
		for j := 0; j <= i; j++ {
			fmt.Print(num, " ")
			num = num * (i - j) / (j + 1)
		}
		fmt.Println()
	}
}

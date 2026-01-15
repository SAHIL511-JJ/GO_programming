// Program to create and print multidimensional slice
package main

import "fmt"

func main() {
	// creating 2D slice (multidimensional)
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println("Multidimensional Slice (Matrix):")
	fmt.Println(matrix)

	fmt.Println("\nPrinting row by row:")
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Print(matrix[i][j], " ")
		}
		fmt.Println()
	}
}

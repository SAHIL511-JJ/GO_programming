// Program to multiply two matrices
package main

import "fmt"

func main() {
	// first matrix 2x3
	matrix1 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	// second matrix 3x2
	matrix2 := [][]int{
		{7, 8},
		{9, 10},
		{11, 12},
	}

	// result matrix 2x2
	rows := len(matrix1)
	cols := len(matrix2[0])
	result := make([][]int, rows)
	for i := range result {
		result[i] = make([]int, cols)
	}

	// matrix multiplication
	for i := 0; i < len(matrix1); i++ {
		for j := 0; j < len(matrix2[0]); j++ {
			for k := 0; k < len(matrix2); k++ {
				result[i][j] += matrix1[i][k] * matrix2[k][j]
			}
		}
	}

	fmt.Println("Matrix 1:")
	for _, row := range matrix1 {
		fmt.Println(row)
	}

	fmt.Println("\nMatrix 2:")
	for _, row := range matrix2 {
		fmt.Println(row)
	}

	fmt.Println("\nResult (Matrix1 x Matrix2):")
	for _, row := range result {
		fmt.Println(row)
	}
}

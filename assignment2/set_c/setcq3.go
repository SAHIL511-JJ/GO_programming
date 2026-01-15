// Program to return multiple values from function
package main

import "fmt"

// function returning multiple values
func getMinMax(numbers []int) (int, int) {
	min := numbers[0]
	max := numbers[0]

	for i := 0; i < len(numbers); i++ {
		if numbers[i] < min {
			min = numbers[i]
		}
		if numbers[i] > max {
			max = numbers[i]
		}
	}
	return min, max
}

func main() {
	nums := []int{5, 2, 9, 1, 7, 3, 8}

	fmt.Println("Numbers:", nums)

	// getting multiple return values
	minimum, maximum := getMinMax(nums)

	fmt.Println("Minimum =", minimum)
	fmt.Println("Maximum =", maximum)
}

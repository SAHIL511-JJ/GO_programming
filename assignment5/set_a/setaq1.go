// Program to calculate sum of squares and cubes using goroutines
package main

import "fmt"

func sumOfSquares(num int, ch chan int) {
	sum := 0
	for num > 0 {
		digit := num % 10
		sum += digit * digit
		num = num / 10
	}
	ch <- sum
}

func sumOfCubes(num int, ch chan int) {
	sum := 0
	for num > 0 {
		digit := num % 10
		sum += digit * digit * digit
		num = num / 10
	}
	ch <- sum
}

func main() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	// creating channels
	sqCh := make(chan int)
	cubeCh := make(chan int)

	// launching goroutines
	go sumOfSquares(num, sqCh)
	go sumOfCubes(num, cubeCh)

	// receiving from channels
	squares := <-sqCh
	cubes := <-cubeCh

	fmt.Println("Sum of squares:", squares)
	fmt.Println("Sum of cubes:", cubes)
	fmt.Println("Final sum:", squares+cubes)
}

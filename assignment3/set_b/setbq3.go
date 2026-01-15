// Program to calculate total and average of student marks
package main

import "fmt"

// student structure
type Student struct {
	RollNo  int
	Name    string
	Mark1   float64
	Mark2   float64
	Mark3   float64
	Total   float64
	Average float64
}

func main() {
	var n int
	fmt.Print("Enter number of students: ")
	fmt.Scan(&n)

	students := make([]Student, n)

	// reading student details
	for i := 0; i < n; i++ {
		fmt.Println("\nStudent", i+1)
		fmt.Print("Roll No: ")
		fmt.Scan(&students[i].RollNo)
		fmt.Print("Name: ")
		fmt.Scan(&students[i].Name)
		fmt.Print("Mark 1: ")
		fmt.Scan(&students[i].Mark1)
		fmt.Print("Mark 2: ")
		fmt.Scan(&students[i].Mark2)
		fmt.Print("Mark 3: ")
		fmt.Scan(&students[i].Mark3)

		// calculating total and average
		students[i].Total = students[i].Mark1 + students[i].Mark2 + students[i].Mark3
		students[i].Average = students[i].Total / 3
	}

	// displaying results
	fmt.Println("\n--- Student Results ---")
	for i := 0; i < n; i++ {
		fmt.Println("\nRoll No:", students[i].RollNo)
		fmt.Println("Name:", students[i].Name)
		fmt.Println("Total:", students[i].Total)
		fmt.Printf("Average: %.2f\n", students[i].Average)
	}
}

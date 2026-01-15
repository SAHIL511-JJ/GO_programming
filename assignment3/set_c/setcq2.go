// Program to find employee with maximum salary
package main

import "fmt"

// employee structure
type Employee struct {
	Eno    int
	Ename  string
	Salary float64
}

func main() {
	var n int
	fmt.Print("Enter number of employees: ")
	fmt.Scan(&n)

	employees := make([]Employee, n)

	// reading employee details
	for i := 0; i < n; i++ {
		fmt.Println("\nEmployee", i+1)
		fmt.Print("Employee No: ")
		fmt.Scan(&employees[i].Eno)
		fmt.Print("Name: ")
		fmt.Scan(&employees[i].Ename)
		fmt.Print("Salary: ")
		fmt.Scan(&employees[i].Salary)
	}

	// finding employee with max salary
	maxIndex := 0
	for i := 1; i < n; i++ {
		if employees[i].Salary > employees[maxIndex].Salary {
			maxIndex = i
		}
	}

	fmt.Println("\n--- Employee with Maximum Salary ---")
	fmt.Println("Employee No:", employees[maxIndex].Eno)
	fmt.Println("Name:", employees[maxIndex].Ename)
	fmt.Println("Salary:", employees[maxIndex].Salary)
}

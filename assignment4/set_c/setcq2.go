// Program to display students sorted by percentage
package main

import (
	"fmt"
	"sort"
)

// student structure
type Student struct {
	RollNo     int
	Name       string
	Percentage float64
}

// for sorting
type ByPercentage []Student

func (a ByPercentage) Len() int           { return len(a) }
func (a ByPercentage) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByPercentage) Less(i, j int) bool { return a[i].Percentage > a[j].Percentage }

// method to display student
func (s Student) Display() {
	fmt.Printf("Roll: %d, Name: %s, Percentage: %.2f%%\n", s.RollNo, s.Name, s.Percentage)
}

func main() {
	students := []Student{
		{101, "Rahul", 75.5},
		{102, "Priya", 88.2},
		{103, "Amit", 65.8},
		{104, "Sneha", 92.1},
		{105, "Karan", 78.9},
	}

	fmt.Println("Before Sorting:")
	for _, s := range students {
		s.Display()
	}

	// sorting in descending order of percentage
	sort.Sort(ByPercentage(students))

	fmt.Println("\nAfter Sorting (Descending by Percentage):")
	for _, s := range students {
		s.Display()
	}
}

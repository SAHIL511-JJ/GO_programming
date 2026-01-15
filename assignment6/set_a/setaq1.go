// Program to sort students by marks using sort package
package main

import (
	"fmt"
	"sort"
)

type Student struct {
	Name  string
	Marks int
}

// implementing sort interface
type ByMarks []Student

func (s ByMarks) Len() int           { return len(s) }
func (s ByMarks) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s ByMarks) Less(i, j int) bool { return s[i].Marks < s[j].Marks }

func main() {
	students := []Student{
		{"Rahul", 85},
		{"Priya", 92},
		{"Amit", 78},
		{"Sneha", 88},
		{"Karan", 65},
	}

	fmt.Println("Before sorting:")
	for _, s := range students {
		fmt.Printf("%s: %d\n", s.Name, s.Marks)
	}

	// sorting by marks
	sort.Sort(ByMarks(students))

	fmt.Println("\nAfter sorting by marks:")
	for _, s := range students {
		fmt.Printf("%s: %d\n", s.Name, s.Marks)
	}
}

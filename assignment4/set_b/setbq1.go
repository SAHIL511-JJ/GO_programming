// Program with pointer receiver for student struct
package main

import "fmt"

// student structure
type Student struct {
	RollNo int
	Name   string
	Marks  int
}

// method with pointer receiver - can modify values
func (s *Student) Show() {
	fmt.Println("Roll No:", s.RollNo)
	fmt.Println("Name:", s.Name)
	fmt.Println("Marks:", s.Marks)
}

// method to update marks using pointer
func (s *Student) UpdateMarks(newMarks int) {
	s.Marks = newMarks
}

func main() {
	student := Student{
		RollNo: 101,
		Name:   "Sahil",
		Marks:  85,
	}

	fmt.Println("Before update:")
	student.Show()

	// updating marks using pointer receiver
	student.UpdateMarks(95)

	fmt.Println("\nAfter update:")
	student.Show()
}

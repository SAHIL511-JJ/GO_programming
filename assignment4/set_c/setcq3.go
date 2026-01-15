// Program to demonstrate embedded interfaces
package main

import "fmt"

// Interface 1
type Reader interface {
	Read()
}

// Interface 2
type Writer interface {
	Write()
}

// Embedded interface
type ReadWriter interface {
	Reader
	Writer
}

// File structure
type File struct {
	Name string
}

// implementing Read method
func (f File) Read() {
	fmt.Println("Reading from file:", f.Name)
}

// implementing Write method
func (f File) Write() {
	fmt.Println("Writing to file:", f.Name)
}

func main() {
	file := File{Name: "test.txt"}

	// using embedded interface
	var rw ReadWriter = file

	fmt.Println("Using Embedded Interface:")
	rw.Read()
	rw.Write()
}

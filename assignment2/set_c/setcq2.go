// Program to create file and use defer statement
package main

import (
	"fmt"
	"os"
)

func main() {
	// creating a file
	file, err := os.Create("hello.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	// defer will close file at end of function
	defer file.Close()
	fmt.Println("File created successfully")

	// writing to file
	file.WriteString("Hello World")
	fmt.Println("Written 'Hello World' to file")

	// defer statement ensures file is closed
	// even if error occurs
	fmt.Println("File will be closed automatically by defer")
}

// Program to append content to file
package main

import (
	"fmt"
	"os"
)

func main() {
	filename := "myfile.txt"

	// creating file with initial content
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	file.WriteString("First line\n")
	file.Close()
	fmt.Println("File created with initial content")

	// opening file in append mode
	file, err = os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	// appending content
	file.WriteString("Second line - appended\n")
	file.WriteString("Third line - appended\n")
	file.Close()
	fmt.Println("Content appended successfully")

	// reading and displaying file content
	content, _ := os.ReadFile(filename)
	fmt.Println("\nFile content:")
	fmt.Println(string(content))
}

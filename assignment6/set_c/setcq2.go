// Program to print file information
package main

import (
	"fmt"
	"os"
)

func main() {
	// creating a test file
	file, _ := os.Create("testfile.txt")
	file.WriteString("This is test content")
	file.Close()

	// getting file info
	fileInfo, err := os.Stat("testfile.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("File Information:")
	fmt.Println("-----------------")
	fmt.Println("Name:", fileInfo.Name())
	fmt.Println("Size:", fileInfo.Size(), "bytes")
	fmt.Println("Mode:", fileInfo.Mode())
	fmt.Println("Modified:", fileInfo.ModTime())
	fmt.Println("Is Directory:", fileInfo.IsDir())
}

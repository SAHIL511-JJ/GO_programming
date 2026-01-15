// Program to read and display book details using structure
package main

import "fmt"

// defining book structure
type Book struct {
	BookID int
	Title  string
	Author string
	Price  float64
}

func main() {
	var n int
	fmt.Print("Enter number of books: ")
	fmt.Scan(&n)

	// creating slice of books
	books := make([]Book, n)

	// reading book details
	for i := 0; i < n; i++ {
		fmt.Println("\nEnter details for book", i+1)
		fmt.Print("Book ID: ")
		fmt.Scan(&books[i].BookID)
		fmt.Print("Title: ")
		fmt.Scan(&books[i].Title)
		fmt.Print("Author: ")
		fmt.Scan(&books[i].Author)
		fmt.Print("Price: ")
		fmt.Scan(&books[i].Price)
	}

	// displaying book details
	fmt.Println("\n--- Book Details ---")
	for i := 0; i < n; i++ {
		fmt.Println("\nBook", i+1)
		fmt.Println("ID:", books[i].BookID)
		fmt.Println("Title:", books[i].Title)
		fmt.Println("Author:", books[i].Author)
		fmt.Println("Price:", books[i].Price)
	}
}

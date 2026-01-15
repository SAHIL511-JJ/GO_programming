// Program with struct author and method show()
package main

import "fmt"

// author structure
type Author struct {
	Name     string
	Branch   string
	Articles int
	Salary   int
}

// method with struct receiver
func (a Author) Show() {
	fmt.Println("Author Name:", a.Name)
	fmt.Println("Branch:", a.Branch)
	fmt.Println("Published Articles:", a.Articles)
	fmt.Println("Salary:", a.Salary)
}

func main() {
	// creating author
	author1 := Author{
		Name:     "Rahul",
		Branch:   "Computer Science",
		Articles: 50,
		Salary:   45000,
	}

	fmt.Println("Author Details:")
	fmt.Println("---------------")
	author1.Show()
}

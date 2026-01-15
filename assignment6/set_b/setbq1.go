// Program to add two integers with unit test
// Save this as setbq1.go
package main

import "fmt"

// Add function to add two integers
func Add(a, b int) int {
	return a + b
}

func main() {
	result := Add(5, 3)
	fmt.Println("5 + 3 =", result)

	result2 := Add(10, 20)
	fmt.Println("10 + 20 =", result2)
}

/*
To test this, create setbq1_test.go with:

package main

import "testing"

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	if result != 5 {
		t.Errorf("Expected 5, got %d", result)
	}
}

Run with: go test
*/

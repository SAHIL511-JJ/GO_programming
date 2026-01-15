// Program to subtract two integers with table test
package main

import "fmt"

// Subtract function
func Subtract(a, b int) int {
	return a - b
}

func main() {
	fmt.Println("10 - 5 =", Subtract(10, 5))
	fmt.Println("20 - 8 =", Subtract(20, 8))
}

/*
Table test in setbq2_test.go:

package main

import "testing"

func TestSubtract(t *testing.T) {
	tests := []struct {
		a, b, expected int
	}{
		{10, 5, 5},
		{20, 8, 12},
		{100, 50, 50},
		{0, 0, 0},
	}

	for _, test := range tests {
		result := Subtract(test.a, test.b)
		if result != test.expected {
			t.Errorf("Subtract(%d, %d) = %d; expected %d",
				test.a, test.b, result, test.expected)
		}
	}
}

Run with: go test
*/

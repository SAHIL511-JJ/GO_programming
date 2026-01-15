// Program to find square with benchmark
package main

import "fmt"

// Square function
func Square(n int) int {
	return n * n
}

func main() {
	fmt.Println("Square of 5:", Square(5))
	fmt.Println("Square of 12:", Square(12))
}

/*
Benchmark in setbq3_test.go:

package main

import "testing"

func TestSquare(t *testing.T) {
	result := Square(5)
	if result != 25 {
		t.Errorf("Expected 25, got %d", result)
	}
}

func BenchmarkSquare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Square(100)
	}
}

Run benchmark with: go test -bench=.
*/

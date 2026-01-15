// Program showing error if variable declared but not used
package main

import "fmt"

func main() {
	var a int = 10
	var b int = 20 // this will cause error if not used
	var c int = 30 // this will cause error if not used

	fmt.Println("Value of a:", a)
	fmt.Println("Value of b:", b)
	fmt.Println("Value of c:", c)

	// Note: If you comment out the Println for b or c
	// the compiler will throw error:
	// "b declared but not used"
	// This is a feature of Go to keep code clean
}

/*
To see the error, comment out one of the Println lines above.
Go compiler will show:
	./setbq3.go:8:6: b declared but not used
*/

// Program to demonstrate type switch
package main

import "fmt"

// function using type switch
func checkType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("Integer:", v)
	case float64:
		fmt.Println("Float:", v)
	case string:
		fmt.Println("String:", v)
	case bool:
		fmt.Println("Boolean:", v)
	default:
		fmt.Println("Unknown type")
	}
}

func main() {
	fmt.Println("Type Switch Demo:")
	fmt.Println("-----------------")

	checkType(42)
	checkType(3.14)
	checkType("Hello Go")
	checkType(true)
	checkType([]int{1, 2, 3})
}

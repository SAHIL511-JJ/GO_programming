// Program to check palindrome using function
package main

import "fmt"

// function to reverse a number
func reverse(n int) int {
	var rev int = 0
	for n > 0 {
		digit := n % 10
		rev = rev*10 + digit
		n = n / 10
	}
	return rev
}

// function to check palindrome
func isPalindrome(n int) bool {
	return n == reverse(n)
}

func main() {
	var num int

	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	if isPalindrome(num) {
		fmt.Println(num, "is a Palindrome")
	} else {
		fmt.Println(num, "is not a Palindrome")
	}
}

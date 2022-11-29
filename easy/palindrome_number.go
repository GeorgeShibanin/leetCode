package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	palindrome := strconv.FormatInt(int64(x), 10)
	for i := 0; i <= len(palindrome)/2; i++ {
		if palindrome[i] != palindrome[len(palindrome)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isPalindrome(-121))
}

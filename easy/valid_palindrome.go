package main

import (
	"fmt"
	"strings"
)

func isNumAlpha(s byte) bool {
	if 48 <= s && s <= 57 || 65 <= s && s <= 90 || 97 <= s && s <= 122 {
		return true
	}
	return false
}

func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	s = strings.ToLower(s)
	fmt.Println(s)
	for i < j {
		fmt.Println(s[i], s[j])
		for i < j && !isNumAlpha(s[i]) {
			i++
		}
		for i < j && !isNumAlpha(s[j]) {
			j--
		}
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
}

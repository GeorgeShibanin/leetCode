package main

func longestPalindrome(s string) string {
	var begining, ending, maxLen int
	for i := 0; i < len(s); i++ {
		leftOdd, rightOdd := adjust(s, i, i)
		leftEven, rightEven := adjust(s, i, i+1)
		if rightOdd-leftOdd > rightEven-leftEven && rightOdd-leftOdd > maxLen {
			maxLen = rightOdd - leftOdd
			begining = leftOdd
			ending = rightOdd
		} else if rightEven-leftEven > maxLen {
			maxLen = rightEven - leftEven
			begining = leftEven
			ending = rightEven
		}
	}
	return s[begining : ending+1]
}

func adjust(s string, left, right int) (int, int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left -= 1
		right += 1
	}
	return left + 1, right - 1
}

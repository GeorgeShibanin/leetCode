package main

import "fmt"

func mySqrt(x int) int {
	left := 0
	right := x
	for left < right {
		middle := (right-left)/2 + left
		if middle*middle <= x && x < (middle+1)*(middle+1) {
			return middle
		} else if x < middle*middle {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return left
}

func main() {
	fmt.Println(mySqrt(1))
}
